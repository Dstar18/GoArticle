package main

import (
	"GoArticle/config"
	"GoArticle/controllers"
	"GoArticle/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// Middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost",
			"http://localhost:8080",
			"http://localhost:8888",
		},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	api := e.Group("/api")
	api.POST("/article", articleController.ArticlesStore)
	api.GET("/article/:limit/:offset", articleController.ArticleGets)
	api.GET("/article/:id", articleController.ArticleGetID)
	api.POST("/article/:id", articleController.ArticleUpdate)
	api.DELETE("/article/:id", articleController.ArticleDelete)

	e.Logger.Fatal(e.Start(":3000"))
}
