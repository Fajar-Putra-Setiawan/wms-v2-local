package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// MustBePositive ensures numeric value positive.
func MustBePositive(value int64) error {
	if value < 0 {
		return fmt.Errorf("value must be positive")
	}
	return nil
}

// CheckStatus validates stock status value.
func CheckStatus(status string) error {
	status = strings.ToLower(strings.TrimSpace(status))
	switch status {
	case "good", "damaged", "unavailable", "in-transit":
		return nil
	default:
		return fmt.Errorf("status must be good/damaged/unavailable/in-transit")
	}
}

// BindJSONOrFail calls c.ShouldBindJSON and return 400 on failure.
func BindJSONOrFail(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		SendValidationError(c, err.Error())
		return false
	}
	return true
}
