package controller

import (
	"net/http"

	"github.com/aryuuu/finkita/internal/domain"
	"github.com/labstack/echo/v4"
)

type accountHandler struct {
	accountService domain.IAccountService
}

func InitAccountHandler(e *echo.Group, accountService domain.IAccountService) {
	h := accountHandler{
		accountService: accountService,
	}

	e.POST("/", h.AddAccount)
	e.GET("/", h.GetAccounts)
    e.GET("/:id", h.GetAccountByID)
	e.PATCH("/:id", h.UpdateAccount)
	e.DELETE("/:id", h.DeleteAccount)
}

func (h accountHandler) AddAccount(c echo.Context) error {
	account := domain.Account{}
	err := c.Bind(&account)

	if err != nil {
		return err
	}

	err = h.accountService.AddAccount(c.Request().Context(), &account)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, account)
}

func (h accountHandler) GetAccounts(c echo.Context) error {
	return c.String(http.StatusOK, "GET /accounts")
}

func (h accountHandler) GetAccountByID(c echo.Context) error {
	return c.String(http.StatusOK, "GET /accounts/:id")
}

func (h accountHandler) UpdateAccount(c echo.Context) error {
	return c.String(http.StatusOK, "PATCH /accounts/:id")
}

func (h accountHandler) DeleteAccount(c echo.Context) error {
	return c.String(http.StatusOK, "DELETE /accounts/:id")
}
