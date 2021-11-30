package routes

import (
	"go-capstone/controllers"

	"github.com/gin-gonic/gin"
)

func RoutesListener() {
	router := gin.Default()
	router.GET("/user/:id", controllers.ReadCsv)
	router.POST("/user", controllers.AddCsv)

	router.Run(":8000")
}
