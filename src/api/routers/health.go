package routers

import (
	"clean-web-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func Health (r *gin.RouterGroup){
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health) 
}