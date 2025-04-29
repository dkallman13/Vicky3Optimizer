package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/dkallman13/Vicky3Optimizer/model"
	"github.com/dkallman13/Vicky3Optimizer/types"
)
func StatePost(c *gin.Context) {
		c.Request.ParseForm()
		idstring := c.Request.Form.Get("stateId")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}
		newstate := types.NewState(id, c.Request.Form.Get("stateName"))
		model.DB.Omit("Provinces").Create(&newstate)
	}