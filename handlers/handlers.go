package handlers

import (
	"log"
	"net/http"

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
	})
}

func ViewPost(c *gin.Context) {
	id := c.Param("id")
	post := repository.GetPost(id)
	log.Print("error hunting", post.Name, post.Content)

	c.HTML(http.StatusOK, "single.html", gin.H{
		"title":       config.SiteConfig.Title,
		"description": config.SiteConfig.Description,
		"post":        repository.GetPost(id),
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
