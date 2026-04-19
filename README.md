> Read on Numoru:
> - https://numoru.com/en/contributions/geo-vs-seo-citaciones-llms
> - https://numoru.com/en/contributions/auditoria-citaciones-llms-go

# geo-audit

Go CLI that audits how often your brand is cited across ChatGPT, Claude, Gemini, Perplexity and Google AI Overview.
Uses **LiteLLM** as unified gateway, **Qdrant** to vectorize responses and detect paraphrased citations, and **Langfuse** for reproducible traces.

Companion to: [Auditar si ChatGPT cita tu marca](https://numoru.com/contribuciones/auditoria-citaciones-llms-go).

## Install

```bash
go install github.com/numoru-ia/geo-audit/cmd/audit@latest
```

## Quick start

```bash
audit -c config.yaml -out results/
open results/report.html
```

## Cost

~2 USD per full run (50 queries × 5 providers).
