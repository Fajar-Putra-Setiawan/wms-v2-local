package repositories

import (
	"time"
	"wms/models"

	"gorm.io/gorm"
)

type ProductMasterRepository interface {
	FindByLocation(location string) ([]models.ProductMaster, error)
	FindByDocumentAndDateRange(documentCode string, from, to time.Time) ([]models.ProductMaster, error)
}

type productMasterRepository struct {
	db *gorm.DB
}

func NewProductMasterRepository(db *gorm.DB) ProductMasterRepository {
	return &productMasterRepository{db: db}
}

func (r *productMasterRepository) FindByLocation(location string) ([]models.ProductMaster, error) {
	var masters []models.ProductMaster
	err := r.db.Where("location = ?", location).Order("created_at DESC").Find(&masters).Error
	return masters, err
}

func (r *productMasterRepository) FindByDocumentAndDateRange(documentCode string, from, to time.Time) ([]models.ProductMaster, error) {
	var masters []models.ProductMaster
	err := r.db.Joins("JOIN product_documents ON product_documents.id = product_masters.document_id::uuid").
		Where("product_documents.code = ? AND product_masters.created_at BETWEEN ? AND ?", documentCode, from, to).
		Order("product_masters.created_at DESC").
		Find(&masters).Error
	return masters, err
}
