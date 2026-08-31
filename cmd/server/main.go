// Command server runs Saksama as a local web app: a dashboard that shows the
// agent working live. It serves the static page in web/ and exposes a small
// JSON/streaming API backed by the SAME frozen engine as cmd/solution — no
// logic is re-implemented here, the handlers just call the agent stages and
// stream their progress.
//
//	go run ./cmd/server                 # listen on :8765
//	go run ./cmd/server -addr :9000     # custom port
//
// A live review (POST /api/review) needs the SAKSAMA_* env (auto-loaded from
// .env). The page itself and /api/meta work with no API key.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/EndPx/saksama/internal/agent"
	"github.com/EndPx/saksama/internal/llm"
	"github.com/EndPx/saksama/internal/memo"
	"github.com/EndPx/saksama/internal/scoring"
	"github.com/EndPx/saksama/internal/statutes"
)

func main() {
	addr := flag.String("addr", ":8765", "listen address")
	webDir := flag.String("web", "web", "directory holding live.html")
	statutesPath := flag.String("statutes", "data/statutes/2026-08.yaml", "statute corpus")
	examplesDir := flag.String("examples", "examples", "sample contracts directory")
	open := flag.Bool("open", true, "open the dashboard in a browser on start")
	flag.Parse()

	_ = llm.LoadEnvFile(".env") // shell env still wins

	corpus, err := agent.LoadCorpus(*statutesPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	srv := &server{corpus: corpus, webDir: *webDir, samples: loadSamples(*examplesDir)}

	mux := http.NewServeMux()
	mux.HandleFunc("/", srv.handleIndex)
	mux.HandleFunc("/api/meta", srv.handleMeta)
	mux.HandleFunc("/api/review", srv.handleReview)

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	url := "http://localhost" + portOf(*addr)
	fmt.Printf("Saksama live  ->  %s\n", url)
	fmt.Println("(a live review needs SAKSAMA_API_KEY in .env; the page and samples do not)")
	if *open {
		go openBrowser(url) // fire-and-forget; the port is already bound above
	}
	if err := http.Serve(ln, mux); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// openBrowser tries to open url in the OS default browser. Best-effort: any
// failure just prints a hint, since the server is already usable at the URL.
func openBrowser(url string) {
	time.Sleep(400 * time.Millisecond) // let Serve start accepting
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "(could not open a browser automatically — open %s yourself)\n", url)
	}
}

type server struct {
	corpus  *statutes.Corpus
	webDir  string
	samples []sample
}

type sample struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Text string `json:"text"`
}

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// no-store: the dashboard is served from disk and may be edited between runs;
	// never let a browser cache a stale copy of the page/JS.
	w.Header().Set("Cache-Control", "no-store")
	http.ServeFile(w, r, filepath.Join(s.webDir, "live.html"))
}

// handleMeta returns the 14 provisions and the sample contracts, so the page
// can draw the coverage map and prefill the editor without an API key.
func (s *server) handleMeta(w http.ResponseWriter, r *http.Request) {
	type prov struct {
		ID, Judul, Dasar, Pasal, Tier, Deteksi, Conf string
	}
	out := struct {
		Statutes []prov   `json:"statutes"`
		Samples  []sample `json:"samples"`
	}{Samples: s.samples}
	for _, p := range s.corpus.Provisions {
		out.Statutes = append(out.Statutes, prov{p.ID, p.Judul, p.DasarHukum, p.Pasal,
			string(p.Tier), string(p.Deteksi), string(p.Confidence)})
	}
	writeJSON(w, out)
}

// evJSON is one streamed progress event (newline-delimited JSON).
type evJSON struct {
	T        string    `json:"t"`                  // "stage" | "result" | "error"
	Key      string    `json:"key,omitempty"`      // stage id
	Status   string    `json:"status,omitempty"`   // "run" | "done"
	Detail   string    `json:"detail,omitempty"`   // human note
	Findings []outFind `json:"findings,omitempty"` // final findings
	Dropped  []outFind `json:"dropped,omitempty"`  // gate-rejected clause findings
	Memo     string    `json:"memo,omitempty"`
	Msg      string    `json:"msg,omitempty"`
}

type outFind struct {
	StatuteID string `json:"statute_id"`
	Judul     string `json:"judul"`
	Dasar     string `json:"dasar"`
	Pasal     string `json:"pasal"`
	Section   string `json:"section"`
	Tier      string `json:"tier"`
	Deteksi   string `json:"deteksi"`
	Kutipan   string `json:"kutipan,omitempty"`
	Deskripsi string `json:"deskripsi,omitempty"`
	Conf      string `json:"conf"`
	Kind      string `json:"kind"` // "clause" | "missing"
	Reason    string `json:"reason,omitempty"`
}

// handleReview runs the real S5 pipeline on the posted contract text and streams
// each stage as it happens: classify -> clause checks -> absence pass -> citation
// gate -> memo. The dashboard renders these events live.
func (s *server) handleReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}
	body, _ := readAll(r, 1<<20)
	text := string(body)
	if strings.TrimSpace(text) == "" {
		http.Error(w, "empty contract", http.StatusBadRequest)
		return
	}

	flush, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/x-ndjson")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("X-Accel-Buffering", "no")

	enc := json.NewEncoder(w)
	emit := func(e evJSON) {
		_ = enc.Encode(e)
		flush.Flush()
	}
	fail := func(msg string) { emit(evJSON{T: "error", Msg: msg}) }

	// Guardrail: only review something that looks like an Indonesian employment
	// contract. If it does not, stop here — no model call, no cost.
	if !looksLikeContract(text) {
		emit(evJSON{T: "notcontract"})
		return
	}

	client, err := llm.New()
	if err != nil {
		fail("No model configured. Set SAKSAMA_API_KEY / SAKSAMA_API_BASE / SAKSAMA_MODEL in .env, then restart the server. " + err.Error())
		return
	}
	a := agent.New(client, s.corpus)
	ctx := r.Context()

	// classify
	typ := "PKWT"
	if strings.Contains(strings.ToUpper(text), "PKWTT") {
		typ = "PKWTT"
	}
	emit(evJSON{T: "stage", Key: "classify", Status: "done", Detail: typ})

	// clause checks
	emit(evJSON{T: "stage", Key: "clause", Status: "run"})
	clause, _, _, err := a.Checklist(ctx, text)
	if err != nil {
		fail("clause checks failed: " + err.Error())
		return
	}
	emit(evJSON{T: "stage", Key: "clause", Status: "done", Detail: fmt.Sprintf("%d flagged", len(clause))})

	// absence pass
	emit(evJSON{T: "stage", Key: "absence", Status: "run"})
	absent, _, _, err := a.Absence(ctx, text)
	if err != nil {
		fail("absence pass failed: " + err.Error())
		return
	}
	emit(evJSON{T: "stage", Key: "absence", Status: "done", Detail: fmt.Sprintf("%d missing", len(absent))})

	// citation gate
	emit(evJSON{T: "stage", Key: "gate", Status: "run"})
	kept, rejected := a.CitationGate(text, clause)
	emit(evJSON{T: "stage", Key: "gate", Status: "done",
		Detail: fmt.Sprintf("%d kept, %d dropped", len(kept), len(rejected))})

	// memo
	emit(evJSON{T: "stage", Key: "memo", Status: "run"})
	findings := append(append([]scoring.Finding{}, kept...), absent...)
	md := memo.Render("kontrak", text, typ, findings, s.corpus)
	emit(evJSON{T: "stage", Key: "memo", Status: "done"})

	// final payload
	out := evJSON{T: "result", Memo: md}
	for _, f := range findings {
		out.Findings = append(out.Findings, s.enrich(f))
	}
	sort.SliceStable(out.Findings, func(i, j int) bool {
		return tierRank(out.Findings[i].Tier) < tierRank(out.Findings[j].Tier)
	})
	for _, rj := range rejected {
		of := s.enrich(rj.Finding)
		of.Reason = rj.Reason
		out.Dropped = append(out.Dropped, of)
	}
	emit(out)
}

func (s *server) enrich(f scoring.Finding) outFind {
	of := outFind{
		StatuteID: f.StatuteID, Section: f.Section,
		Tier: string(f.Tier), Deteksi: string(f.Deteksi),
		Kutipan: f.Kutipan, Deskripsi: f.Deskripsi,
	}
	if p, ok := s.corpus.Get(f.StatuteID); ok {
		of.Judul, of.Dasar, of.Pasal = p.Judul, p.DasarHukum, p.Pasal
		of.Conf = string(p.Confidence)
		if of.Tier == "" {
			of.Tier = string(p.Tier)
		}
		if p.Deteksi == statutes.DeteksiTidakAdaKlausa {
			of.Kind = "missing"
		} else {
			of.Kind = "clause"
		}
	}
	return of
}

func tierRank(t string) int {
	switch t {
	case string(statutes.TierBatalDemiHukum):
		return 0
	case string(statutes.TierSanksiAdministratif):
		return 1
	case string(statutes.TierMelanggarTanpaSanksi):
		return 2
	case string(statutes.TierPedomanKebijakan):
		return 3
	}
	return 4
}

func loadSamples(dir string) []sample {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var out []sample
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".txt") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		id := strings.TrimSuffix(e.Name(), ".txt")
		typ := "PKWT"
		if strings.Contains(strings.ToUpper(string(b)), "PKWTT") {
			typ = "PKWTT"
		}
		out = append(out, sample{ID: id, Name: prettyName(id), Type: typ, Text: string(b)})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func prettyName(id string) string {
	id = strings.ReplaceAll(id, "-", " ")
	parts := strings.Fields(id)
	for i, p := range parts {
		switch strings.ToLower(p) {
		case "pkwt", "pkwtt":
			parts[i] = strings.ToUpper(p)
		default:
			if len(p) > 0 {
				parts[i] = strings.ToUpper(p[:1]) + p[1:]
			}
		}
	}
	return strings.Join(parts, " ")
}

// looksLikeContract is a cheap, deterministic guard (no model call) that decides
// whether the pasted text plausibly is an Indonesian employment contract. It
// requires a strong employment-agreement signal plus several supporting terms,
// so obvious non-contracts (prose, code, a shopping list) are rejected before
// the pipeline runs.
func looksLikeContract(text string) bool {
	t := strings.ToLower(text)
	if len(strings.TrimSpace(t)) < 200 {
		return false
	}
	strong := strings.Contains(t, "pkwt") ||
		strings.Contains(t, "pkwtt") ||
		strings.Contains(t, "kontrak kerja") ||
		(strings.Contains(t, "perjanjian") && strings.Contains(t, "kerja"))
	if !strong {
		return false
	}
	signals := 0
	for _, kw := range []string{
		"pihak", "pasal", "pekerja", "karyawan", "pengusaha", "pemberi kerja",
		"upah", "gaji", "masa kerja", "jangka waktu", "hubungan kerja", "jabatan",
	} {
		if strings.Contains(t, kw) {
			signals++
		}
	}
	return signals >= 3
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func readAll(r *http.Request, max int64) ([]byte, error) {
	defer r.Body.Close()
	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 4096)
	var total int64
	for {
		n, err := r.Body.Read(tmp)
		if n > 0 {
			total += int64(n)
			if total > max {
				return buf, nil
			}
			buf = append(buf, tmp[:n]...)
		}
		if err != nil {
			return buf, nil
		}
	}
}

func portOf(addr string) string {
	if strings.HasPrefix(addr, ":") {
		return addr
	}
	if i := strings.LastIndex(addr, ":"); i >= 0 {
		return addr[i:]
	}
	return ":" + addr
}
