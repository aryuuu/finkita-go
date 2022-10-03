package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/aryuuu/finkita/domain"
	"github.com/labstack/echo/v4"
)

type mutationHandler struct {
	mutationService domain.IMutationService
}

func InitMutationHandler(e *echo.Group, mutationService domain.IMutationService) {
	h := mutationHandler{
		mutationService: mutationService,
	}

	e.GET("/:id", h.GetMutationByID)
	e.GET("", h.GetMutations)
}

func (h mutationHandler) GetMutations(c echo.Context) error {
	email := c.Request().Header.Get("x-user-email")
	mutations, err := h.mutationService.GetMutationsByEmail(c.Request().Context(), email)
	if err != nil {
		log.Printf("error fetching mutations: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, mutations)
}

func (h mutationHandler) GetMutationByID(c echo.Context) error {
	id := c.Param("id")
	account, err := h.mutationService.GetMutationByID(context.Background(), id)
	if err != nil {
		log.Printf("error fetching mutation by id: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, account)
}
