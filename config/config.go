package config

import (
	"github.com/BurntSushi/toml"
	"github.com/niliaranet/forum/models"
	"log"
	"os"
)

var SiteConfig models.SiteConfig

func LoadSiteConfig() {
	var conf models.SiteConfig

	configContent, err := os.ReadFile("site.toml")
	if err != nil {
		log.Panic(err)
	}

	toml.Decode(string(configContent), &conf)
	SiteConfig = conf
}
