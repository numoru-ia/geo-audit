package runner

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/numoru-ia/geo-audit/internal/config"
	"github.com/numoru-ia/geo-audit/internal/detector"
	"github.com/numoru-ia/geo-audit/internal/providers"
)

type Result struct {
	Query     string             `json:"query"`
	Provider  string             `json:"provider"`
	Response  string             `json:"response"`
	Latency   time.Duration      `json:"latency_ms"`
	Citations detector.Citations `json:"citations"`
	Err       string             `json:"err,omitempty"`
}

type Runner struct {
	cfg         *config.Config
	concurrency int
	providers   *providers.Registry
	detector    *detector.Detector
}

func New(cfg *config.Config, concurrency int) *Runner {
	apiKey := os.Getenv("LITELLM_MASTER_KEY")
	return &Runner{
		cfg:         cfg,
		concurrency: concurrency,
		providers:   providers.New(cfg.LiteLLMBaseURL, apiKey),
		detector:    detector.New(),
	}
}

type task struct {
	q        string
	provider config.Provider
}

func (r *Runner) Run(ctx context.Context) ([]Result, error) {
	tasks := make(chan task)
	results := make(chan Result)
	var wg sync.WaitGroup

	for i := 0; i < r.concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range tasks {
				start := time.Now()
				resp, err := r.providers.Ask(ctx, t.provider, t.q)
				lat := time.Since(start)
				res := Result{
					Query:    t.q,
					Provider: t.provider.ID,
					Response: resp,
					Latency:  lat,
				}
				if err != nil {
					res.Err = err.Error()
				} else {
					res.Citations = r.detector.Find(t.q, resp, r.cfg.Brand, r.cfg.Domains, r.cfg.Competitors)
				}
				results <- res
			}
		}()
	}

	go func() {
		defer close(tasks)
		for _, q := range r.cfg.Queries {
			for _, p := range r.cfg.Providers {
				select {
				case <-ctx.Done():
					return
				case tasks <- task{q: q, provider: p}:
				}
			}
		}
	}()

	go func() { wg.Wait(); close(results) }()

	var out []Result
	for r := range results {
		out = append(out, r)
	}
	return out, nil
}
