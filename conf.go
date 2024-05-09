package main

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Enable       bool   `yaml:"enable_casdoor"`
	Endpoint     string `yaml:"endpoint"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Organization string `yaml:"organization"`
	Application  string `yaml:"application"`
	FrontendURL  string `yaml:"frontend_url"`
	RedirectPath string `yaml:"redirect_path"`
}

type Config struct {
	Certificate string       `yaml:"certificate"`
	Server      ServerConfig `yaml:"server"`
}

var GlobalConfig *Config

func LoadConfig(configPath string) error {
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	GlobalConfig = &cfg

	return nil
}
