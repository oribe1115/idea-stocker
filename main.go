package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e.Start(":" + port)
}
