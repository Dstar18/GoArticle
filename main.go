package main

import (
	"GoArticle/config"
	"GoArticle/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// initialize env, DB & migrate
	config.LoadEnv()
	db := config.ConnectDB()

	// initialize echo
	e := echo.New()

	routes.SetupRoutes(e, db)

	e.Logger.Fatal(e.Start(":3000"))
}
