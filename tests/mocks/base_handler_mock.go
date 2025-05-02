package mocks

import (
	"DesignRestaurantManagementSystem/interfaces"

	"github.com/gin-gonic/gin"
)

// MockBaseHandler is a mock implementation of BaseHandlerInterface
type MockBaseHandler struct {
	// SuccessResponseFunc will be called when SuccessResponse is called
	SuccessResponseFunc func(c *gin.Context, status int, data interface{})

	// ErrorResponseFunc will be called when ErrorResponse is called
	ErrorResponseFunc func(c *gin.Context, status int, message string)

	// ValidationErrorFunc will be called when ValidationError is called
	ValidationErrorFunc func(c *gin.Context, err error)

	// InternalServerErrorFunc will be called when InternalServerError is called
	InternalServerErrorFunc func(c *gin.Context, err error)
}

// Ensure MockBaseHandler implements BaseHandlerInterface
var _ interfaces.BaseHandlerInterface = (*MockBaseHandler)(nil)

func (m *MockBaseHandler) SuccessResponse(c *gin.Context, status int, data interface{}) {
	if m.SuccessResponseFunc != nil {
		m.SuccessResponseFunc(c, status, data)
		return
	}
	c.JSON(status, gin.H{
		"status": "success",
		"data":   data,
	})
}

func (m *MockBaseHandler) ErrorResponse(c *gin.Context, status int, message string) {
	if m.ErrorResponseFunc != nil {
		m.ErrorResponseFunc(c, status, message)
		return
	}
	c.JSON(status, gin.H{
		"status":  "error",
		"message": message,
	})
}

func (m *MockBaseHandler) ValidationError(c *gin.Context, err error) {
	if m.ValidationErrorFunc != nil {
		m.ValidationErrorFunc(c, err)
		return
	}
	c.JSON(400, gin.H{
		"status":  "error",
		"message": "Validation failed",
		"errors":  err.Error(),
	})
}

func (m *MockBaseHandler) InternalServerError(c *gin.Context, err error) {
	if m.InternalServerErrorFunc != nil {
		m.InternalServerErrorFunc(c, err)
		return
	}
	c.JSON(500, gin.H{
		"status":  "error",
		"message": "Internal server error",
	})
}
