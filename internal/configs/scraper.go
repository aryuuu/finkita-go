package configs

import "os"

type scraper struct {
	MyUsername string
	MyPassword string
}

func setupScraper() *scraper {
	return &scraper{
		MyUsername: os.Getenv("MY_USERNAME"),
		MyPassword: os.Getenv("MY_PASSWORD"),
	}
}

