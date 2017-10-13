package router

import (
	"github.com/PrinceNorin/bakanovels/config"
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once
var router *gin.Engine

func Get() *gin.Engine {
	once.Do(func() {
		if config.Get().Environment == "production" {
			gin.SetMode(gin.ReleaseMode)
		}
		router = gin.New()
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	})
	return router
}
