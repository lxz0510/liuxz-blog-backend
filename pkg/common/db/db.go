package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"liuxz-blog-backend/pkg/common/models"
)

func Init(url string) *gorm.DB {
	//
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&models.Article{})
	return db
}
