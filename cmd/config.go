package cmd

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Port    string `env:"PORT" envDefault:"3030"`
	DBHost  string `env:"DB_HOST" envDefault:"db"`
DBPort  string `env:"DB_PORT" envDefault:"3306"`
DBUser  string `env:"DB_USER " envDefault:"root"`
DBPass  string `env:"DB_PASSWORD" envDefault:"password"`
DBNAME  string `env:"DB_DATABASE" envDefault:"db"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	log.Printf("%+v\n", cfg)
	return cfg, nil
}

func (c *Config) GetDB() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBNAME)
	return sqlx.Open("mysql", dsn)
}