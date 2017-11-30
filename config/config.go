package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config contains app configuration includes Server, RethinkDB
type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	RethinkDB struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		DB       string `json:"db"`
	} `json:"rethinkdb"`
	Options struct {
		Prefix string `json:"prefix"`
	} `json:"options"`
}

func FromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
