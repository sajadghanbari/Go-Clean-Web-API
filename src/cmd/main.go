package main

import (
	"clean-web-api/api"
	"clean-web-api/config"
	"clean-web-api/data/cache"
	"clean-web-api/data/db"
	"clean-web-api/pkg/logging"

)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(),nil)
	}
	defer cache.CloseRedis()
	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(),nil)
	}
	defer db.CloseDb()
	api.InitServer(cfg)

}
