package repositories

import (
	"wms/models"

	"gorm.io/gorm"
)

type ProductMasterRepository interface {
	FindByLocation(location string) ([]models.ProductMaster, error)
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
