package configs

import "os"

type postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func setupPostgres() *postgres {
	return &postgres{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		Username: os.Getenv("PG_USERNAME"),
		Password: os.Getenv("PG_PASSWORD"),
		Database: os.Getenv("PG_DATABASE"),
	}
}
