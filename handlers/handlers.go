package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/niliaranet/forum/config"
	"github.com/niliaranet/forum/models"
	"github.com/niliaranet/forum/repository"
)

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       config.SiteConfig.Title,
		"description": config.SiteConfig.Description,
		"posts":       repository.GetPosts(),
		"timeParse":   time.Parse,
	})
}

func NewPostForm(c *gin.Context) {
	c.HTML(http.StatusOK, "new-post.html", gin.H{
		"title":       config.SiteConfig.Title,
		"description": config.SiteConfig.Description,
	})
}

func CreatePost(c *gin.Context) {
	newPost := models.Post{
		Name:    c.PostForm("name"),
		Content: c.PostForm("content"),
	}

	repository.CreatePost(newPost)
	c.Redirect(http.StatusFound, "/")
}
