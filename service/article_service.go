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

func (s *ArticleService) GetsArticle(limit, offset int) ([]models.Posts, error) {
	var articles []models.Posts
	result := s.DB.Limit(limit).Offset(offset).Find(&articles)
	return articles, result.Error
}

func (s *ArticleService) GetIdArticle(id uint) (models.Posts, error) {
	var article models.Posts
	err := s.DB.First(&article, id).Error
	return article, err
}
