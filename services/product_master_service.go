package services

import (
	"wms/models"
	"wms/repositories"
)

type ProductMasterService interface {
	GetByLocation(location string) ([]models.ProductMaster, error)
}

type productMasterService struct {
	repo repositories.ProductMasterRepository
}

func NewProductMasterService(repo repositories.ProductMasterRepository) ProductMasterService {
	return &productMasterService{repo: repo}
}

func (s *productMasterService) GetByLocation(location string) ([]models.ProductMaster, error) {
	return s.repo.FindByLocation(location)
}
