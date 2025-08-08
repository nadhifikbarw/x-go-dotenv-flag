package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/nadhifikbarw/x-go-dotenv-flag/pkg/config"
)

func main() {
	// Configure -env flag with env_file path support
	// adjust flag parsing logic based on project
	// this one is when using `flag` stdlib
	var envFile string
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

	// See config loaded
	config, err := config.Load()
	if err != nil {
		log.Fatalf("config: %+v", err)
	}

	fmt.Printf("config: %+v", config)

}
