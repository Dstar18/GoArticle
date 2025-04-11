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

func (s *ArticleService) UpdateArticle(data *models.Posts) error {
	return s.DB.Save(data).Error
}

func (s *ArticleService) DeleteArticle(id uint) error {
	return s.DB.Delete(&models.Posts{}, id).Error
}

func (s *ArticleService) GetArticlesByStatus(status string) ([]models.Posts, error) {
	var article []models.Posts
	result := s.DB.Where("status = ?", status).Find(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return article, nil
}
