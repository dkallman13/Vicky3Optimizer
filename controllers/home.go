package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/dkallman13/Vicky3Optimizer/model"
	"net/http"
)
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title":     "home",
		"savefiles": model.SaveFiles,
	})
}