package controller

import (
	"log"
	"net/http"

	"github.com/aryuuu/finkita/domain"
	"github.com/labstack/echo/v4"
)

type accountHandler struct {
	accountService domain.IAccountService
}

func InitAccountHandler(e *echo.Group, accountService domain.IAccountService) {
	h := accountHandler{
		accountService: accountService,
	}

	e.GET("/:id", h.GetAccountByID)
	e.PATCH("/:id", h.UpdateAccountByID)
	e.DELETE("/:id", h.DeleteAccountByID)
	e.POST("", h.AddAccount)
	e.GET("", h.GetAccounts)
}

func (h accountHandler) AddAccount(c echo.Context) error {
	account := domain.Account{}
	err := c.Bind(&account)

	if err != nil {
		return err
	}

	email := c.Request().Header.Get("x-user-email")
	account.Email = email

	err = h.accountService.AddAccount(c.Request().Context(), &account)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, account)
}

func (h accountHandler) GetAccounts(c echo.Context) error {
	log.Println("GET /accounts")
	email := c.Request().Header.Get("x-user-email")
	accounts, err := h.accountService.GetAccountsByEmail(c.Request().Context(), email)

	if err != nil {
		log.Printf("error in get accounts controller, propagating error: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, accounts)
}

func (h accountHandler) GetAccountByID(c echo.Context) error {
	id := c.Param("id")
	account, err := h.accountService.GetAccountByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, account)
}

func (h accountHandler) UpdateAccountByID(c echo.Context) error {
	account := domain.Account{}
	err := c.Bind(&account)

	if err != nil {
		return err
	}

	err = h.accountService.UpdateAccountByID(c.Request().Context(), "", &account)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, account)
}

func (h accountHandler) DeleteAccountByID(c echo.Context) error {
	id := c.Param("id")
	err := h.accountService.DeleteAccount(c.Request().Context(), id)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
