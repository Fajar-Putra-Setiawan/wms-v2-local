package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PaginationParams holds paging query values.
type PaginationParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

// PaginationMeta details response metadata.
type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
	TotalRows   int `json:"total_rows"`
	TotalPages  int `json:"total_pages"`
}

// ParsePagination reads from query params.
func ParsePagination(c *gin.Context, defaultLimit int) PaginationParams {
	page := 1
	limit := defaultLimit

	if p := c.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}

	offset := (page - 1) * limit
	return PaginationParams{Page: page, Limit: limit, Offset: offset}
}

// BuildPaginationMeta builds response metadata.
func BuildPaginationMeta(totalRows, page, limit int) PaginationMeta {
	totalPages := 1
	if limit > 0 {
		totalPages = int(math.Ceil(float64(totalRows) / float64(limit)))
	}
	if totalPages < 1 {
		totalPages = 1
	}
	return PaginationMeta{CurrentPage: page, PageSize: limit, TotalRows: totalRows, TotalPages: totalPages}
}
