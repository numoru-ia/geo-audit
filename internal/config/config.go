package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Provider struct {
	ID    string `yaml:"id"`
	Model string `yaml:"model"`
}

type Competitor struct {
	Brand   string   `yaml:"brand"`
	Domains []string `yaml:"domains"`
}

type Config struct {
	Brand       string       `yaml:"brand"`
	Domains     []string     `yaml:"domains"`
	Competitors []Competitor `yaml:"competitors"`
	Queries     []string     `yaml:"queries"`
	Providers   []Provider   `yaml:"providers"`

	LiteLLMBaseURL string `yaml:"litellm_base_url"`
	QdrantURL      string `yaml:"qdrant_url"`
	LangfuseURL    string `yaml:"langfuse_url"`
}

func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Config
	return &c, yaml.Unmarshal(b, &c)
}
