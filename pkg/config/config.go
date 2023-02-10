package config

import (
	"encoding/json"
	"io/ioutil"
	"xr-central/pkg/db"
)

type Config struct {
	Server Server    `json:"Server"`
	DB     db.Config `json:"DB"`
	GCP    GCP       `json:"GCP"`
	Hello  Hello     `json:"Hello"`
}

type Server struct {
	Port string `json:"Port"`
}

type GCP struct {
	ProjectID string `json:"ProjectID"`
	LogName   string `json:"LogName"`
}

type Hello struct {
	Show string `json:"Show"`
}

func New(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
