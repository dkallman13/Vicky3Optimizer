package controllers

import (
	"log"
	"os"

	"github.com/dkallman13/Vicky3Optimizer/model"
	"github.com/gin-gonic/gin"
)

func StateFileU(c *gin.Context) {
	c.Request.ParseMultipartForm(8000000) //8MB limit
	c.Request.ParseForm()
	stateFile, err := c.FormFile("statefile")
	if err != nil {
		log.Println("error here")
		log.Fatal(err)
	}
	log.Println(stateFile.Filename)
	c.SaveUploadedFile(stateFile, "./statefiles/"+stateFile.Filename)

	openedFile, err2 := os.Open("./statefiles/"+stateFile.Filename)
	if err2 != nil {
		log.Println("error here2")
		log.Fatal(err2)
	}
	defer openedFile.Close()

	model.TokenizeStateFile(openedFile)
	
}
