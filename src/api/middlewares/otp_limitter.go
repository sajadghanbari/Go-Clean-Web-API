package middlewares

import (
	"clean-web-api/api/helper"
	"clean-web-api/config"
	limiter "clean-web-api/pkg/limiter"
	"errors"
	"net/http"
	"time"

	"golang.org/x/time/rate"
	"github.com/gin-gonic/gin"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIPRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(c *gin.Context) {
		limiter := limiter.GetLimiter(c.Request.RemoteAddr)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, helper.OtpLimiterError, errors.New("not allowed")))
			c.Abort()
		} else {
			c.Next()
		}
	}
}