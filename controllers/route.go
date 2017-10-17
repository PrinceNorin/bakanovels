package controllers

import (
	"github.com/PrinceNorin/bakanovels/controllers/novels"
	"github.com/PrinceNorin/bakanovels/controllers/router"
	"github.com/PrinceNorin/bakanovels/controllers/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var APIRouter *gin.Engine

func init() {
	am := buildAuthMiddleware()

	r := router.Get()

	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	conf.AddAllowMethods("OPTIONS")
	api := r.Group("/api", cors.New(conf))
	{
		api.POST("/login", am.LoginHandler)
		api.OPTIONS("/login", am.LoginHandler)

		api.POST("/register", userController.UserRegisterHandler)
		api.OPTIONS("/register", userController.UserRegisterHandler)
	}

	v1 := api.Group("/v1", am.MiddlewareFunc())
	{
		v1.POST("/novels", novelController.CreateNovelHandler)
		v1.OPTIONS("/novels", novelController.CreateNovelHandler)
	}

	APIRouter = r
}
