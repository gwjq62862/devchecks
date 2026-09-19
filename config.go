package goterminal

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type DevConfig struct {
	Envs  []string `yaml:"envs"`
	Hosts []string `yaml:"hosts"`
}

func LoadFile(filename string) ([]byte, error) {
	if filename == "" {
		filename = "devcheck.yaml"
	}
	byte, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	return byte, nil
}

func LoadConfig(filename string) (*DevConfig, error) {
	byte, err := LoadFile(filename)
	if err != nil {
		return nil, err
	}
	//emty object instace
	var cfg DevConfig

	err = yaml.Unmarshal(byte, &cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid YAML syntax: %w", err)
	}
	if len(cfg.Envs) == 0 && len(cfg.Hosts) == 0 {
		return nil, fmt.Errorf("config file is empty or missing required sections")
	}

	return &cfg, nil
}
