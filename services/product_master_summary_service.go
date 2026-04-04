package services

import (
	"time"
	"wms/repositories"
)

type ProductMasterSummary struct {
	TotalPieces      int     `json:"total_pieces"`
	TotalHargaAsal   float64 `json:"total_harga_asal"`
	TotalHargaGudang float64 `json:"total_harga_gudang"`
}

type ProductMasterSummaryService interface {
	GetSummary(from, to time.Time) (ProductMasterSummary, error)
}

type productMasterSummaryService struct {
	repo repositories.ProductMasterRepository
}

func NewProductMasterSummaryService(repo repositories.ProductMasterRepository) ProductMasterSummaryService {
	return &productMasterSummaryService{repo: repo}
}

func (s *productMasterSummaryService) GetSummary(from, to time.Time) (ProductMasterSummary, error) {
	masters, err := s.repo.FindByDocumentAndDateRange("INBOUND_MANUAL", from, to)
	if err != nil {
		return ProductMasterSummary{}, err
	}
	summary := ProductMasterSummary{}
	for _, m := range masters {
		summary.TotalPieces += m.Item
		summary.TotalHargaAsal += float64(m.Item) * m.Price
		summary.TotalHargaGudang += float64(m.Item) * m.PriceWarehouse
	}
	return summary, nil
}
