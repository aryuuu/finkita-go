package main

import (
	"fmt"

	"github.com/aryuuu/finkita/internal/controller"
	"github.com/aryuuu/finkita/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Printf("Hello World\n")
	e := echo.New()
	healtcheckGroup := e.Group("/healthcheck")
	accountGroup := e.Group("/accounts")
	accountService := service.NewAccountService()

	// TODO: setup repositories
	// TODO: setup services
	// TODO: setup controllers
	controller.InitHealthCheckHandler(healtcheckGroup)
	controller.InitAccountHandler(accountGroup, accountService)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World")
	// })
	e.Logger.Fatal(e.Start(":8080"))
}
