package services

import (
	"fmt"
	"math/rand"
	"time"
	"wms/models"

	"gorm.io/gorm"
)

type InboundService interface {
	InboundManual(req models.InboundRequest, db *gorm.DB) (pending models.ProductPending, master models.ProductMaster, err error)
}

type inboundService struct{}

func NewInboundService() InboundService {
	return &inboundService{}
}

func (s *inboundService) InboundManual(req models.InboundRequest, db *gorm.DB) (models.ProductPending, models.ProductMaster, error) {
	// Logic generate barcode
	barcode := generateUniqueBarcode()
	barcodeWarehouse := generateUniqueBarcode()

	// Logic dokumen: cari/buat dokumen manual
	doc, err := getOrCreateManualDocument(db)
	if err != nil {
		return models.ProductPending{}, models.ProductMaster{}, err
	}

	// Logic BE: tentukan category_id/sticker_id otomatis dan PriceWarehouse
	var categoryID, stickerID, typeID string
	var priceWarehouse float64 = req.Price
	var master models.ProductMaster
	if req.Price >= 100000 {
		if req.CategoryID != nil {
			categoryID = *req.CategoryID
			// Ambil diskon kategori
			var category models.Category
			if err := db.Where("id = ?", categoryID).First(&category).Error; err == nil && category.Discount != nil {
				discount := float64(*category.Discount)
				priceWarehouse = req.Price * (1 - discount/100)
			}
		}
		stickerID = ""
		typeID = "categories"

		master = models.ProductMaster{
			DocumentID:       doc.ID.String(),
			Barcode:          barcode,
			BarcodeWarehouse: barcodeWarehouse,
			Name:             req.Name,
			NameWarehouse:    "Manual",
			Item:             req.Item,
			Price:            req.Price,
			PriceWarehouse:   priceWarehouse,
			CategoryID:       categoryID,
			StickerID:        stickerID,
			TypeID:           typeID,
			Location:         "staging_reguler",
			TypeOut:          "cargo",
		}
	} else {
		if req.StickerID != nil {
			stickerID = *req.StickerID
			// Cari sticker sesuai range harga
			var sticker models.Sticker
			if err := db.Where("id = ?", stickerID).First(&sticker).Error; err == nil && sticker.MinPrice != nil && sticker.MaxPrice != nil {
				if req.Price >= float64(*sticker.MinPrice) && req.Price <= float64(*sticker.MaxPrice) && sticker.FixedPrice != nil {
					priceWarehouse = float64(*sticker.FixedPrice)
				}
			}
		} else {
			// Jika stickerID tidak ada, cari sticker yang cocok dengan range harga
			var sticker models.Sticker
			if err := db.Where("min_price <= ? AND max_price >= ?", req.Price, req.Price).First(&sticker).Error; err == nil && sticker.FixedPrice != nil {
				stickerID = sticker.ID.String()
				priceWarehouse = float64(*sticker.FixedPrice)
			}
		}
		categoryID = ""
		typeID = "sticker"

		master = models.ProductMaster{
			DocumentID:       doc.ID.String(),
			Barcode:          barcode,
			BarcodeWarehouse: barcodeWarehouse,
			Name:             req.Name,
			NameWarehouse:    "Manual",
			Item:             req.Item,
			Price:            req.Price,
			PriceWarehouse:   priceWarehouse,
			CategoryID:       categoryID,
			StickerID:        stickerID,
			TypeID:           typeID,
			Location:         "staging_sticker",
			TypeOut:          "cargo",
		}
	}

	// Insert ke ProductPending
	pending := models.ProductPending{
		DocumentID: doc.ID.String(),
		Barcode:    barcode,
		Name:       req.Name,
		Item:       req.Item,
		Price:      req.Price,
		Status:     req.Status, // default status valid
	}
	if err := db.Create(&pending).Error; err != nil {
		return pending, master, err
	}

	if err := db.Create(&master).Error; err != nil {
		return pending, master, err
	}

	return pending, master, nil
}

// Helper: generate barcode dan dokumen, bisa diambil dari controller lama
func generateUniqueBarcode() string {
	t := time.Now().UnixNano()
	r := rand.Intn(100000)
	return fmt.Sprintf("BC-%d-%d", t, r)
}

func getOrCreateManualDocument(db *gorm.DB) (models.ProductDocument, error) {
	var doc models.ProductDocument
	err := db.Where("code = ?", "INBOUND_MANUAL").First(&doc).Error
	if err == gorm.ErrRecordNotFound {
		doc = models.ProductDocument{
			Code:     "INBOUND_MANUAL",
			FileName: "INBOUND_MANUAL",
			Type:     "manual",
			Status:   "progress",
		}
		if err := db.Create(&doc).Error; err != nil {
			return doc, err
		}
		return doc, nil
	}
	return doc, err
}
