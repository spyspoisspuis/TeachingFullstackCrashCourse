package handlers

import "github.com/gin-gonic/gin"

type Handler interface {
	InsertUser(c *gin.Context)
}