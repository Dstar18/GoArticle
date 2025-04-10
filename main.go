package main

import (
	"GoArticle/config"
	"GoArticle/controllers"
	"GoArticle/service"

	"github.com/labstack/echo/v4"
)

func main() {
	// initialize env, DB
	config.LoadEnv()
	dbConn := config.ConnectDB()

	// Init Service
	articleService := service.NewArticleService(dbConn)
	// Init Controller
	articleController := controllers.NewArticleController(articleService)

	// Routes
	e := echo.New()
	api := e.Group("/api")
	api.POST("/article", articleController.ArticlesStore)
	api.GET("/article/:limit/:offset", articleController.ArticleGets)
	api.GET("/article/:id", articleController.ArticleGetID)

	e.Logger.Fatal(e.Start(":3000"))
}
