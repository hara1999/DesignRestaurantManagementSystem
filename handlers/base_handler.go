package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	logger *slog.Logger
}

func NewBaseHandler(logger *slog.Logger) *BaseHandler {
	return &BaseHandler{
		logger: logger,
	}
}

func (h *BaseHandler) SuccessResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{
		"status": "success",
		"data":   data,
	})
}

func (h *BaseHandler) ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  "error",
		"message": message,
	})
}

func (h *BaseHandler) ValidationError(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"status":  "error",
		"message": "Validation failed",
		"errors":  err.Error(),
	})
}

func (h *BaseHandler) InternalServerError(c *gin.Context, err error) {
	h.logger.Error("Internal server error", "error", err)
	c.JSON(500, gin.H{
		"status":  "error",
		"message": "Internal server error",
	})
}
