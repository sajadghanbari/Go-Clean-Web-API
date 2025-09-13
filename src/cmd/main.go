package main

import (
	"clean-web-api/api"
	"clean-web-api/config"
	"clean-web-api/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)

}