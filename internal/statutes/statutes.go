// Package statutes loads and represents the Saksama legal corpus: the fixed set
// of Indonesian employment-law provisions that findings are matched against.
package statutes

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Tier is the severity class of a provision.
type Tier string

const (
	TierBatalDemiHukum       Tier = "batal_demi_hukum"
	TierSanksiAdministratif  Tier = "sanksi_administratif"
	TierMelanggarTanpaSanksi Tier = "melanggar_tanpa_sanksi"
	TierPedomanKebijakan     Tier = "pedoman_kebijakan"
)

// Detection is how a violation of a provision is detected.
type Detection string

const (
	DeteksiAdaKlausa      Detection = "ada_klausa"
	DeteksiTidakAdaKlausa Detection = "tidak_ada_klausa"
	DeteksiKonteks        Detection = "konteks"
)

// Confidence records how well a provision was verified.
type Confidence string

const (
	ConfidenceA Confidence = "A" // verified against primary text
	ConfidenceB Confidence = "B" // consistent secondary sources only
)

// Provision is one entry in the legal corpus. Field tags map to the Indonesian
// YAML keys; the statutory text itself stays Indonesian.
type Provision struct {
	ID         string     `yaml:"id"`
	DasarHukum string     `yaml:"dasar_hukum"`
	Pasal      string     `yaml:"pasal"`
	Tier       Tier       `yaml:"tier"`
	Confidence Confidence `yaml:"confidence"`
	Judul      string     `yaml:"judul"`
	Ringkasan  string     `yaml:"ringkasan"`
	Deteksi    Detection  `yaml:"deteksi"`
}

// Corpus is the full set of provisions, indexed by ID.
type Corpus struct {
	Provisions []Provision
	byID       map[string]Provision
}

// Get returns the provision with the given id, if present.
func (c *Corpus) Get(id string) (Provision, bool) {
	p, ok := c.byID[id]
	return p, ok
}

// IDs returns the valid provision ids in file order.
func (c *Corpus) IDs() []string {
	ids := make([]string, len(c.Provisions))
	for i, p := range c.Provisions {
		ids[i] = p.ID
	}
	return ids
}

var validTiers = map[Tier]bool{
	TierBatalDemiHukum:       true,
	TierSanksiAdministratif:  true,
	TierMelanggarTanpaSanksi: true,
	TierPedomanKebijakan:     true,
}

var validDeteksi = map[Detection]bool{
	DeteksiAdaKlausa:      true,
	DeteksiTidakAdaKlausa: true,
	DeteksiKonteks:        true,
}

var validConfidence = map[Confidence]bool{ConfidenceA: true, ConfidenceB: true}

// Load reads and validates the corpus at path. It fails on an empty corpus,
// empty or duplicate ids, or any invalid tier/deteksi/confidence value.
func Load(path string) (*Corpus, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read statutes: %w", err)
	}
	var provs []Provision
	if err := yaml.Unmarshal(data, &provs); err != nil {
		return nil, fmt.Errorf("parse statutes: %w", err)
	}
	if len(provs) == 0 {
		return nil, fmt.Errorf("statutes: corpus is empty")
	}
	byID := make(map[string]Provision, len(provs))
	for i, p := range provs {
		if p.ID == "" {
			return nil, fmt.Errorf("statutes: entry %d has empty id", i)
		}
		if _, dup := byID[p.ID]; dup {
			return nil, fmt.Errorf("statutes: duplicate id %q", p.ID)
		}
		if !validTiers[p.Tier] {
			return nil, fmt.Errorf("statutes: %s has invalid tier %q", p.ID, p.Tier)
		}
		if !validDeteksi[p.Deteksi] {
			return nil, fmt.Errorf("statutes: %s has invalid deteksi %q", p.ID, p.Deteksi)
		}
		if !validConfidence[p.Confidence] {
			return nil, fmt.Errorf("statutes: %s has invalid or empty confidence %q", p.ID, p.Confidence)
		}
		byID[p.ID] = p
	}
	return &Corpus{Provisions: provs, byID: byID}, nil
}
