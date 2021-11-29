package routes

import (
	"go-capstone/controllers"

	"github.com/gin-gonic/gin"
)

func RoutesListener() {
	router := gin.Default()
	router.GET("/user/:id", controllers.ReadCsv)

	router.Run("localhost:8000")
}
