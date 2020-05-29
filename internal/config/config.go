package config

import "github.com/caarlos0/env/v6"

type Options struct {
	Environment   string `env:"ENV"`
	ParserAddress string `env:"PARSER_ADDRESS"`
	SearchAddress string `env:"SEARCH_ADDRESS"`
	DBUser        string `env:"DB_USER"`
	DBName        string `env:"DB_NAME"`
	DBPassword    string `env:"DB_PASSWORD"`
	DBHostname    string `env:"DB_HOST"`
	DBPort        string `env:"DB_PORT"`
}

func NewConfig() (Options, error) {
	opts := Options{}

	if err := env.Parse(&opts); err != nil {
		return Options{}, err
	}

	return opts, nil
}
