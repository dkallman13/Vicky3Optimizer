package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/dkallman13/Vicky3Optimizer/model"
	"github.com/dkallman13/Vicky3Optimizer/types"
	"net/http"

)
func StateGet(c *gin.Context){
	var results []types.State
	model.DB.Table("states").Select("*").Find(&results)
	c.HTML(http.StatusOK, "/db/state", gin.H{
		"title":  "states",
		"states": results,
	})
}