package middleware

import (
	"errors"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			idTokenHeader := c.Request().Header.Get("id-token")
			if idTokenHeader == "" {
				return errors.New("no auth token provided")
			}

			audience := ""
			payload, err := idtoken.Validate(c.Request().Context(), idTokenHeader, audience)
			if err != nil {
				return err
			}

			userEmail := payload.Claims["email"].(string)
			c.Request().Header.Set("x-user-email", userEmail)
			next(c)
			return nil
		}
	}
}
