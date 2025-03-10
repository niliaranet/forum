package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/niliaranet/forum/handlers"
)

var router = gin.Default()

func Run() {
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", handlers.MainPage)
	router.GET("/view/:id", handlers.ViewPost)
	router.GET("/newPost", handlers.NewPostForm)
	router.POST("/newPost", handlers.CreatePost)
	router.Run("localhost:8080")
}
