package configs

import "os"

type scraper struct {
	MyUsername        string
	MyPassword        string
	SeleniumServerURL string
	BNIMobileWebURL   string
}

func setupScraper() *scraper {
	return &scraper{
		MyUsername:        os.Getenv("MY_USERNAME"),
		MyPassword:        os.Getenv("MY_PASSWORD"),
		SeleniumServerURL: os.Getenv("SELENIUM_SERVER_URL"),
		BNIMobileWebURL:   os.Getenv("BNI_MOBILE_WEB_URL"),
	}
}
