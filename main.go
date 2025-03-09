package main

import (
	"github.com/niliaranet/forum/config"
	"github.com/niliaranet/forum/repository"
	"github.com/niliaranet/forum/routes"
)

func main() {
	config.LoadSiteConfig()
	repository.Load()

	routes.Run()
}
