package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ArticleValidate struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}

func ArticlesStore(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var article ArticleValidate

		// Check request body
		if err := c.Bind(&article); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Invalid request body",
			})
		}

		// Validation struct
		validate := validator.New()
		if err := validate.Struct(article); err != nil {
			errors := make(map[string]string)
			for _, err := range err.(validator.ValidationErrors) {
				errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
			}
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": errors,
			})
		}

		// return success
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Created successfully",
		})
	}
}
