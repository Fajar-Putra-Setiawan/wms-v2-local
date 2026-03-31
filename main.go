package main

import (
	"github.com/gin-gonic/gin"
	"wms/utils"
)

type StockRequest struct {
	SKU       string `json:"sku" binding:"required"`
	Quantity  int64  `json:"quantity" binding:"required,min=1"`
	Status    string `json:"status" binding:"required"`
	UnitPrice string `json:"unit_price" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		utils.SendSuccess(c, gin.H{"message": "pong"}, "OK")
	})

	r.POST("/stock/in", func(c *gin.Context) {
		var req StockRequest
		if !utils.BindJSONOrFail(c, &req) {
			return
		}

		if err := utils.CheckStatus(req.Status); err != nil {
			utils.SendError(c, 400, err.Error())
			return
		}

		price, err := utils.ParseCurrency(req.UnitPrice)
		if err != nil {
			utils.SendError(c, 400, "unit_price invalid")
			return
		}

		formattedPrice := utils.FormatCurrency(price, "Rp", 0)
		available := utils.CalcAvailableStock(req.Quantity, 0, 0)

		res := gin.H{
			"sku":             req.SKU,
			"quantity_in":     req.Quantity,
			"status":          req.Status,
			"unit_price":      formattedPrice,
			"available_stock": available,
		}

		utils.SendSuccess(c, res, "stock in recorded")
	})

	r.Run(":8080")
}

