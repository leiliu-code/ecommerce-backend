package main

import (
	"github.com/gin-gonic/gin"

	"github.com/leiliu-code/ecommerce-backend/internal/config"
	"github.com/leiliu-code/ecommerce-backend/internal/router"
)

func main() {
	cfg := config.Load()
	r := gin.Default()
	router.Setup(r)
	r.Run(cfg.Port)
}
