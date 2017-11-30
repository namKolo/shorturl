package main

import (
	"flag"
	"log"

	config "github.com/namKolo/shorturl/config"
)

func main() {
	configPath := flag.String("config", "./config/config.json", "path of the config file")
	flag.Parse()
	appConfig, err := config.FromFile(*configPath)

	if err != nil {
		log.Fatal(err)
	}

	app, err := NewApp(appConfig)
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
