package main

import (
	"GoArticle/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// connect DB
	config.InitDB()

	// initialize echo
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
