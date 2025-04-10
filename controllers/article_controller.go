package controllers

import (
	"GoArticle/models"
	"GoArticle/service"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	articleService *service.ArticleService
}

func NewArticleController(articleService *service.ArticleService) *ArticleController {
	return &ArticleController{articleService: articleService}
}

type ArticleValidate struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}

func (h *ArticleController) ArticlesStore(c echo.Context) error {
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

	param := models.Posts{
		Title:        article.Title,
		Content:      article.Content,
		Category:     article.Category,
		Status:       article.Status,
		Created_date: time.Now(),
		Updated_date: time.Now(),
	}

	// create to service
	if err := h.articleService.CreateArticle(&param); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Created successfully",
		"data":    nil,
	})
}
