package repositories

import (
	"wms/models"

	"gorm.io/gorm"
)

type ProductDocumentRepository interface {
	FindAll() ([]models.ProductDocument, error)
}

type productDocumentRepository struct {
	db *gorm.DB
}

func NewProductDocumentRepository(db *gorm.DB) ProductDocumentRepository {
	return &productDocumentRepository{db: db}
}

func (r *productDocumentRepository) FindAll() ([]models.ProductDocument, error) {
	var docs []models.ProductDocument
	err := r.db.Order("created_at DESC").Find(&docs).Error
	return docs, err
}
