package controllers

import (
	"GoArticle/models"
	"GoArticle/service"
	"net/http"
	"strconv"
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

func (h *ArticleController) ArticleGets(c echo.Context) error {
	limitParam := c.Param("limit")
	offsetParam := c.Param("offset")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid limit",
		})
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid offset",
		})
	}

	// gets to service
	articles, err := h.articleService.GetsArticle(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// manual mapping
	var articleResponses []ArticleValidate
	for _, article := range articles {
		articleResponses = append(articleResponses, ArticleValidate{
			Title:    article.Title,
			Content:  article.Content,
			Category: article.Category,
			Status:   article.Status,
		})
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Article successfully",
		"data":    articleResponses,
	})
}

func (h *ArticleController) ArticleGetID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid id",
		})
	}

	//  get to service
	article, err := h.articleService.GetIdArticle(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Data not found",
		})
	}

	// manual mapping
	var articleResponses ArticleValidate
	articleResponses = ArticleValidate{
		Title:    article.Title,
		Content:  article.Content,
		Category: article.Category,
		Status:   article.Status,
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Article successfully",
		"data":    articleResponses,
	})
}

func (h *ArticleController) ArticleDelete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid id",
		})
	}

	// check data by ID
	_, errCheck := h.articleService.GetIdArticle(uint(id))
	if errCheck != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Data not found",
		})
	}

	// delete to service
	if err := h.articleService.DeleteArticle(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Delete successfully",
		"data":    nil,
	})
}
