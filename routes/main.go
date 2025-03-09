package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() {
	router.LoadHTMLGlob("templates/*")

	router.GET("/", returnHello)
	router.Run("localhost:8080")
}

func returnHello(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"status": "success!",
	})
}

