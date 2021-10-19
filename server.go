package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// create echo instance
	e := echo.New()

	// routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo!")
	})

	// raise server
	e.Logger.Fatal(e.Start(":1323"))
}
