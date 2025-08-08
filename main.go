package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/nadhifikbarw/x-go-dotenv-flag/pkg/config"
)

func main() {
	var envFile string
	// Configure -env flag with optional env_file path support
	flag.BoolFunc("env", "Load .env file", func(s string) error {
		if s == "true" {
			s = ".env"
		}

		envFile = s
		return nil
	})
	flag.Parse()

	if envFile != "" {
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("env: %v", err)
		} else {
			log.Printf("env %q loaded", envFile)
		}
	}

	config, err := config.Load()
	if err != nil {
		log.Fatalf("config: %+v", err)
	}

	fmt.Printf("config: %+v", config)

}
