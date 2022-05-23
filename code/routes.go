package main

import (
	"github.com/ahloul/loyalty-reports/purchase/delivery/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitializeRouts initialize endpoints repos and usecases
func InitializeRouts(router *gin.Engine) {

	h := http.PurchaseHandler{}
	h.Index(router)
}

// InitializeConfig initialize default configurations
func InitializeConfig(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

}
