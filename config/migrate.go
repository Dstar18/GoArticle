package config

import "GoArticle/entity"

func InitMigrate() {
	DB.AutoMigrate(&entity.Posts{})
}
