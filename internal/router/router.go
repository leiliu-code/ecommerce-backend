package router

import (
	"github.com/gin-gonic/gin"

	"github.com/leiliu-code/ecommerce-backend/internal/handler"
)

func Setup(r *gin.Engine) {
	r.GET("/healthz", handler.Health)
}
