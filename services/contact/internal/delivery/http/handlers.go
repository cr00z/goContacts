package http

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

func (d *ContactDelivery) InitRoutes() *gin.Engine {
	mode := gin.DebugMode
	if viper.GetBool("IS_PROD") &&
		strings.ToUpper(strings.TrimSpace(viper.GetString("LOG_LEVEL"))) != "DEBUG" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	router := gin.New()

	return router
}
