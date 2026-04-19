package report

import (
	"encoding/json"
	"html/template"
	"os"

	"github.com/numoru-ia/geo-audit/internal/runner"
)

const htmlTpl = `<!DOCTYPE html>
<html lang="es">
<head><meta charset="utf-8"><title>GEO audit report</title>
<style>body{font-family:system-ui;margin:2em;}table{border-collapse:collapse}td,th{border:1px solid #ccc;padding:4px 8px;}</style>
</head>
<body>
<h1>GEO audit report</h1>
<p>Queries: {{len .}}</p>
<table>
<tr><th>Query</th><th>Provider</th><th>Latency (ms)</th><th>Score</th><th>Literal</th><th>Competitors</th></tr>
{{range .}}<tr>
  <td>{{.Query}}</td>
  <td>{{.Provider}}</td>
  <td>{{.Latency}}</td>
  <td>{{printf "%.2f" .Citations.Score}}</td>
  <td>{{.Citations.LiteralMatch}}</td>
  <td>{{range .Citations.CompetitorsMatched}}{{.}} {{end}}</td>
</tr>{{end}}
</table>
</body></html>`

func WriteJSON(path string, results []runner.Result) error {
	b, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

func WriteHTML(path string, results []runner.Result) error {
	tpl, err := template.New("r").Parse(htmlTpl)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return tpl.Execute(f, results)
}
