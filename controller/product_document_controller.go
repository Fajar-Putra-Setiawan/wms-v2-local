package controller

import (
	"net/http"
	"wms/services"
	"wms/utils"

	"github.com/gin-gonic/gin"
)

type ProductDocumentController struct {
	service services.ProductDocumentService
}

func NewProductDocumentController(service services.ProductDocumentService) *ProductDocumentController {
	return &ProductDocumentController{service: service}
}

func (ctl *ProductDocumentController) ListDocuments(c *gin.Context) {
	docs, err := ctl.service.ListDocuments()
	if err != nil {
		utils.SendError(c, 500, err.Error())
		return
	}
	utils.SendSuccess(c, docs, "List product documents", nil, http.StatusOK)
}
