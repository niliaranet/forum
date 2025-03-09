package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/niliaranet/forum/handlers"
)

var router = gin.Default()

func Run() {
	fmt.Println("running")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.MainPage)
	router.Run("localhost:8080")
}
