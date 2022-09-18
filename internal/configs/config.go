package configs

var Service *service
var Postgres *postgres
var Account *account

func init() {
	Service = setupService()
	Postgres = setupPostgres()
	Account = setupAccount()
}
