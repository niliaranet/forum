package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/niliaranet/forum/config"
	"github.com/niliaranet/forum/repository"
)

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": config.SiteConfig.Title,
		"description": config.SiteConfig.Description,
		"posts": repository.GetPosts(),
	})
}
