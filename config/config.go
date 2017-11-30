package config

import (
	"encoding/json"
	"io/ioutil"
)

// RethinkDB contains db configuration
type RethinkDB struct {
	Host string `json:"host"`
	DB   string `json:"db"`
}

// App contains app configuration includes Server, RethinkDB
type App struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	RethinkDB RethinkDB `json:"rethinkdb"`
	Options   struct {
		Prefix string `json:"prefix"`
	} `json:"options"`
}

// FromFile reads configuration from provided path and return App Configuration
func FromFile(path string) (*App, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg App
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
