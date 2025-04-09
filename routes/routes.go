package routes

import (
	"GoArticle/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	e.POST("/article", controllers.ArticlesStore(db))
}
