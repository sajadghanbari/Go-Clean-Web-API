package main

import (
	"clean-web-api/api"
	"clean-web-api/config"
	"clean-web-api/data/cache"
	"clean-web-api/data/db"
	"log"
)
// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()
	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDb()
	api.InitServer(cfg)

}
