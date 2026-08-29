// Package memo renders the final human-facing artifact: an Indonesian triage
// memo for the worker. It must not read as machine output — no JSON, no
// technical labels, no emoji.
package memo

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/EndPx/saksama/internal/scoring"
	"github.com/EndPx/saksama/internal/statutes"
)

// tierOrder is the display order and Indonesian heading for each tier.
var tierOrder = []struct {
	Tier    statutes.Tier
	Heading string
}{
	{statutes.TierBatalDemiHukum, "Batal demi hukum"},
	{statutes.TierSanksiAdministratif, "Dapat dikenai sanksi administratif"},
	{statutes.TierMelanggarTanpaSanksi, "Melanggar ketentuan (tanpa jalur sanksi)"},
	{statutes.TierPedomanKebijakan, "Pedoman kebijakan (daya paksa terbatas)"},
}

var (
	companyRe  = regexp.MustCompile(`\*\*((?:PT|CV)\s[^*]+?)\*\*`)
	positionRe = regexp.MustCompile(`jabatan\s+([A-Z][^.,\n(]+?)(?:\s+dan ditempatkan|\.|,|\n)`)
)

func firstMatch(re *regexp.Regexp, s, fallback string) string {
	if m := re.FindStringSubmatch(s); len(m) > 1 {
		return strings.TrimSpace(m[1])
	}
	return fallback
}

// Render builds the markdown memo for one reviewed contract.
func Render(contractID, contractText, jenis string, findings []scoring.Finding, corpus *statutes.Corpus) string {
	company := firstMatch(companyRe, contractText, "perusahaan")
	position := firstMatch(positionRe, contractText, "posisi yang ditawarkan")

	// Partition findings.
	present := map[statutes.Tier][]scoring.Finding{}
	var absent []scoring.Finding
	tierCount := map[statutes.Tier]int{}
	var serious bool
	for _, f := range findings {
		tierCount[f.Tier]++
		if f.Tier == statutes.TierBatalDemiHukum || f.Tier == statutes.TierSanksiAdministratif {
			serious = true
		}
		if f.Deteksi == statutes.DeteksiTidakAdaKlausa {
			absent = append(absent, f)
			continue
		}
		present[f.Tier] = append(present[f.Tier], f)
	}

	var b strings.Builder
	fmt.Fprintf(&b, "# Telaah Kontrak Kerja — %s di %s\n\n", position, company)
	fmt.Fprintf(&b, "Tanggal penilaian: %s. Penilaian ini didasarkan pada hukum ketenagakerjaan Indonesia yang berlaku per Agustus 2026. Perlu dicatat bahwa terdapat rancangan undang-undang pelindungan ketenagakerjaan yang sedang dibahas, sehingga ketentuan dapat berubah.\n\n", time.Now().Format("2 January 2006"))

	// Executive summary (<= 5 sentences).
	b.WriteString("## Ringkasan\n\n")
	if len(findings) == 0 {
		fmt.Fprintf(&b, "Dari penelaahan, tidak ditemukan hal yang perlu ditandai pada kontrak %s ini. Anda tetap disarankan membaca ulang seluruh isi kontrak sebelum menandatangani.\n\n", contractID)
	} else {
		parts := []string{}
		for _, to := range tierOrder {
			if n := tierCount[to.Tier]; n > 0 {
				parts = append(parts, fmt.Sprintf("%d hal %s", n, strings.ToLower(to.Heading)))
			}
		}
		fmt.Fprintf(&b, "Ditemukan %d hal yang perlu Anda perhatikan: %s. ", len(findings), strings.Join(parts, "; "))
		if serious {
			b.WriteString("Beberapa di antaranya menyangkut ketentuan yang berat, sehingga sebaiknya Anda mengklarifikasi ke bagian personalia sebelum menandatangani.\n\n")
		} else {
			b.WriteString("Sebagian besar bersifat perlu diklarifikasi, bukan larangan menandatangani.\n\n")
		}
	}

	// Findings by tier (clause-based).
	b.WriteString("## Temuan\n\n")
	anyPresent := false
	for _, to := range tierOrder {
		fs := present[to.Tier]
		if len(fs) == 0 {
			continue
		}
		anyPresent = true
		fmt.Fprintf(&b, "### %s\n\n", to.Heading)
		for _, f := range fs {
			writeFinding(&b, f, corpus)
		}
	}
	if !anyPresent {
		b.WriteString("Tidak ada klausa bermasalah yang tertulis secara eksplisit.\n\n")
	}

	// Provisions not found.
	b.WriteString("## Ketentuan wajib yang tidak ditemukan\n\n")
	if len(absent) == 0 {
		b.WriteString("Tidak ada pelindungan wajib yang tampak hilang dari kontrak ini.\n\n")
	} else {
		for _, f := range absent {
			p, _ := corpus.Get(f.StatuteID)
			fmt.Fprintf(&b, "- **%s.** %s (dasar: %s %s).\n", p.Judul, oneLine(f.Deskripsi, p.Ringkasan), p.DasarHukum, p.Pasal)
		}
		b.WriteString("\n")
	}

	// Questions for HR.
	if len(findings) > 0 {
		b.WriteString("## Pertanyaan untuk bagian personalia\n\n")
		for _, f := range findings {
			p, _ := corpus.Get(f.StatuteID)
			fmt.Fprintf(&b, "- %s\n", hrQuestion(p))
		}
		b.WriteString("\n")
	}

	// When to seek help.
	if serious {
		b.WriteString("## Kapan sebaiknya mencari bantuan\n\n")
		b.WriteString("Karena terdapat temuan pada tingkat yang berat, sebaiknya Anda berkonsultasi dengan dinas ketenagakerjaan setempat, serikat pekerja, atau lembaga bantuan hukum sebelum menandatangani kontrak ini.\n\n")
	}

	// Mandatory closing.
	b.WriteString("## Catatan penting\n\n")
	b.WriteString("Dokumen ini adalah alat bantu telaah awal yang dihasilkan secara otomatis, bukan nasihat hukum. Dokumen ini tidak menggantikan pertimbangan penasihat hukum yang berkualifikasi, dan keputusan akhir sepenuhnya ada pada Anda.\n")
	return b.String()
}

func writeFinding(b *strings.Builder, f scoring.Finding, corpus *statutes.Corpus) {
	p, _ := corpus.Get(f.StatuteID)
	fmt.Fprintf(b, "**%s** (%s)\n\n", p.Judul, f.Section)
	if f.Kutipan != "" {
		fmt.Fprintf(b, "> %s\n\n", strings.TrimSpace(f.Kutipan))
	}
	fmt.Fprintf(b, "Dasar hukum: %s %s. %s Tingkat keyakinan penilaian: %s.\n\n",
		p.DasarHukum, p.Pasal, oneLine(f.Deskripsi, p.Ringkasan), p.Confidence)
}

func oneLine(primary, fallback string) string {
	s := strings.TrimSpace(primary)
	if s == "" {
		s = fallback
	}
	s = strings.Join(strings.Fields(s), " ")
	if !strings.HasSuffix(s, ".") {
		s += "."
	}
	return s
}

func hrQuestion(p statutes.Provision) string {
	return fmt.Sprintf("Mengenai %s: bisakah dijelaskan bagaimana ketentuan ini diterapkan, dan apakah dapat disesuaikan agar sejalan dengan %s %s?",
		strings.ToLower(p.Judul), p.DasarHukum, p.Pasal)
}
