package detector

import (
	"strings"

	"github.com/numoru-ia/geo-audit/internal/config"
)

type Citations struct {
	LiteralMatch       bool     `json:"literal_match"`
	DomainMatches      []string `json:"domain_matches,omitempty"`
	SemanticMatch      bool     `json:"semantic_match"`
	CompetitorsMatched []string `json:"competitors_matched,omitempty"`
	Score              float64  `json:"score"`
}

type Detector struct{}

func New() *Detector { return &Detector{} }

func (d *Detector) Find(query, response, brand string, domains []string, competitors []config.Competitor) Citations {
	var c Citations
	low := strings.ToLower(response)
	if strings.Contains(low, strings.ToLower(brand)) {
		c.LiteralMatch = true
	}
	for _, dom := range domains {
		if strings.Contains(low, strings.ToLower(dom)) {
			c.DomainMatches = append(c.DomainMatches, dom)
		}
	}
	for _, comp := range competitors {
		if strings.Contains(low, strings.ToLower(comp.Brand)) {
			c.CompetitorsMatched = append(c.CompetitorsMatched, comp.Brand)
		}
	}
	// score: 0.6 literal + 0.4 domain; adjust per preference
	if c.LiteralMatch {
		c.Score += 0.6
	}
	if len(c.DomainMatches) > 0 {
		c.Score += 0.4
	}
	return c
}
