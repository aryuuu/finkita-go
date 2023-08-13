package configs

var Service *service
var Postgres *postgres
var Account *account
var Scraper *scraper

func init() {
	Service = setupService()
	Postgres = setupPostgres()
	Account = setupAccount()
	Scraper = setupScraper()
}
