package main

import (
	"flag"
	"log"

	"github.com/viogami/FavAni/config"
	"github.com/viogami/FavAni/server"
)

var (
	appConfig = flag.String("config", "config/app.yaml", "application config path")
)

func main() {
	conf, err := config.ConfigParse(appConfig)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	s, err := server.New(conf)
	if err != nil {
		log.Fatalf("Init server failed: %v", err)
	}

	if err := s.Run(); err != nil {
		log.Fatalf("Run server failed: %v", err)
	}
}
