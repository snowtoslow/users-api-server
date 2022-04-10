package main

import (
	"log"
	"users-api-server/cmd"
)

func main() {
	if err := cmd.Run(8080);err != nil {
		log.Fatal("Failed to run users-api")
	}
}
