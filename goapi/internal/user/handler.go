package user

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/tquocminh17/goapi/pkg/models"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}
