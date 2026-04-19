package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/numoru-ia/geo-audit/internal/config"
	"github.com/numoru-ia/geo-audit/internal/runner"
	"github.com/numoru-ia/geo-audit/internal/report"
)

func main() {
	cfgPath := flag.String("c", "config.yaml", "config path")
	out := flag.String("out", "results", "output dir")
	concurrency := flag.Int("n", 50, "parallel workers")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	cfg, err := config.Load(*cfgPath)
	if err != nil {
		logger.Error("load config", "err", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	r := runner.New(cfg, *concurrency)
	results, err := r.Run(ctx)
	if err != nil {
		logger.Error("run", "err", err)
		os.Exit(1)
	}
	if err := os.MkdirAll(*out, 0o755); err != nil {
		logger.Error("mkdir", "err", err)
		os.Exit(1)
	}
	if err := report.WriteJSON(*out+"/results.json", results); err != nil {
		logger.Error("json", "err", err)
	}
	if err := report.WriteHTML(*out+"/report.html", results); err != nil {
		logger.Error("html", "err", err)
	}
	logger.Info("done", "queries", len(results))
}
