package main

import (
	"database/sql"
	"fmt"

	"github.com/aryuuu/finkita/internal/configs"
	"github.com/aryuuu/finkita/internal/controller"
	"github.com/aryuuu/finkita/internal/middleware"
	"github.com/aryuuu/finkita/internal/repositories"
	"github.com/aryuuu/finkita/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Printf("Hello World\n")
	e := echo.New()
	e.Use(middleware.ErrorLogger())
	healtcheckGroup := e.Group("/healthcheck")
	apiV1Group := e.Group("/api/v1")
	accountGroup := apiV1Group.Group("/accounts")
	mutationGroup := apiV1Group.Group("/mutations")

	db := createDBConnection()
	defer db.Close()
	accountRepo := repositories.NewAccountRepo(db)
	mutationRepo := repositories.NewMutationRepo(db)

	accountService := service.NewAccountService(accountRepo)
	mutationService := service.NewMutationService(mutationRepo)

	controller.InitHealthCheckHandler(healtcheckGroup)
	controller.InitAccountHandler(accountGroup, accountService)
	controller.InitMutationHandler(mutationGroup, mutationService)

	e.Logger.Fatal(e.Start(":8080"))
}

func createDBConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.Postgres.Host, configs.Postgres.Port, configs.Postgres.Username, configs.Postgres.Password, configs.Postgres.Database)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("failed to create connection to postgres db: %v\n", err)
		panic(err)
	}

	return db
}
