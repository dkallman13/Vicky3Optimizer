package model

import(
	"log"
	"github.com/joho/godotenv"
)

func GetEnvVars(){
	err:= godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
}