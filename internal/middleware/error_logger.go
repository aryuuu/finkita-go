package middleware

import (
	"log"

	echo "github.com/labstack/echo/v4"
)

func ErrorLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			err = next(c)
			if err == nil {
				return
			}

			resCtx := map[string]interface{}{
				"method":          c.Request().Method,
				"uri":             c.Request().RequestURI,
				"err":             err,
				"request_context": c.Get("request_context"),
			}

			log.Printf("error: %v", resCtx)

			return err
		}
	}
}
