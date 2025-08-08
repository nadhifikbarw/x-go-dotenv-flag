// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Nadhif Ikbar Wibowo

package config

import (
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	DBHost     string `envconfig:"HOST" required:"true"`
	DBPort     Port   `envconfig:"PORT" required:"true"`
	DBUser     string `envconfig:"USER" required:"true"`
	DBPassword string `envconfig:"PASSWORD"`
	DBDatabase string `envconfig:"DATABASE" required:"true"`
}

func Load() (*DBConfig, error) {
	return LoadPrefix("DB")
}

func LoadPrefix(prefix string) (*DBConfig, error) {
	var config DBConfig
	err := envconfig.Process(prefix, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
