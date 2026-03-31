package utils

import "github.com/gin-gonic/gin"

// APIResponse built standard success payload.
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// APIError represent error payload.
type APIError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// SendSuccess returns 200 response.
func SendSuccess(c *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "success"
	}
	c.JSON(200, APIResponse{Status: "success", Message: message, Data: data})
}

// SendError returns error response.
func SendError(c *gin.Context, code int, message string) {
	if code == 0 {
		code = 400
	}
	if message == "" {
		message = "error"
	}
	c.JSON(code, APIError{Status: "error", Message: message, Code: code})
}

// SendValidationError returns 422.
func SendValidationError(c *gin.Context, errors interface{}) {
	c.JSON(422, gin.H{"status": "fail", "errors": errors})
}
