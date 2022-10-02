package middleware

import (
	"errors"
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// TODO: write the actual auth implementation
			// TODO: get auth token from header
			idTokenHeader := c.Request().Header.Get("id-token")
			if idTokenHeader == "" {
				// log.Println("no auth token provided")
				return errors.New("no auth token provided")
				// return next(c)
			}

			// TODO: validate auth token
			audience := ""
			payload, err := idtoken.Validate(c.Request().Context(), idTokenHeader, audience)
			if err != nil {
				return err
			}

			// TODO: write auth related data if auth passed
			userEmail := payload.Claims["email"].(string)
			c.Request().Header.Set("x-user-email", userEmail)
			log.Printf("userEmail: %s", userEmail)
			// TODO: return err if auth not passed
			next(c)
			return nil
		}
	}
}
