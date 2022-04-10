package main

import (
	"flag"
	"log"
	"users-api-server/cmd"
	"users-api-server/config"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "./config/default.config.json", "config file (default is $HOME/./config/default.config.json)")
	flag.Parse()
}

func main() {
	cfg, err := config.ParseConfigFile(cfgPath)
	if err != nil {
		log.Fatal("Failed to parse config file by specified path:", err)
	}

	if err = cmd.Run(cfg); err != nil {
		log.Fatal("Failed to run users-api")
	}

	log.Printf("The users-api-server is up and running on port: %s with sqllite database name: %s", cfg.AppPort, cfg.DatabaseName)
}
