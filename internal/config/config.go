package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env               string `env:"env"`
	APP_ADDR          uint   `env:"APP_ADDR"`
	Service_DB_DSN    string `env:"SERVICE_DB_DSN"`
	Postgres_DB       string `env:"POSTGRES_DB"`
	Postgres_user     string `env:"POSTGRES_USER"`
	Postgres_password string `env:"POSTGRES_PASSWORD"`
	PGData            string `env:"PGDATA"`
}

func MustReadEnv() *Config {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
