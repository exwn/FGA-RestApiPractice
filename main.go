package main

import (
	"practice_01/controllers"

	"./config"
	"./controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/persons", inDB.GetPersons)
}
