package routers

import (
	"clean-web-api/api/handlers"
	"clean-web-api/api/middlewares"
	"clean-web-api/config"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/send-otp", h.SendOtp)
	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
}