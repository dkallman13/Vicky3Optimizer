package controllers

import (
	"net/http"
	"strings"
	"github.com/dkallman13/Vicky3Optimizer/model"
	"github.com/gin-gonic/gin"
)

func Save(router *gin.Engine){
	for _, savFile := range model.SaveFiles {
		var savRouteBuilder strings.Builder
		savRouteBuilder.WriteString("save/")
		savRouteBuilder.WriteString(savFile)
		router.GET(savRouteBuilder.String(), func(c *gin.Context) {
			c.HTML(http.StatusOK, "/save", gin.H{
				"title":     savFile,
				"firstline": model.TokenizeSave(savFile),
			})
		})
	}
}
