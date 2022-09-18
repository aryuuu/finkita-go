package configs

import "os"

type service struct {
	Name string
	Port string
}

func setupService() *service {
	return &service{
		Name: os.Getenv("SERVICE_NAME"),
		Port: os.Getenv("SERVICE_PORT"),
	}
}
