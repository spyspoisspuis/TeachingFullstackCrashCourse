package handlers

import (
	"backend/users/models"
	"backend/users/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerImpl struct {
	usecase usecases.Usecase
}

func NewHandlerImpl(usecase usecases.Usecase) Handler {
	return &HandlerImpl{
		usecase: usecase,
	}
}

func (h *HandlerImpl) InsertUser(c *gin.Context) {
	var userModel models.Users
	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.InsertUser(&userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User has been created"})
}

func (h *HandlerImpl) GetAllUser(c *gin.Context) {
	users, err := h.usecase.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
