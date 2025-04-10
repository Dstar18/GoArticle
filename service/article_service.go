package service

import (
	"GoArticle/models"

	"gorm.io/gorm"
)

type ArticleService struct {
	DB *gorm.DB
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{DB: db}
}

func (s *ArticleService) CreateArticle(data *models.Posts) error {
	return s.DB.Create(data).Error
}
