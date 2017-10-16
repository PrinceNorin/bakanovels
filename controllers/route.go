package controllers

import (
	"github.com/PrinceNorin/bakanovels/controllers/router"
	_ "github.com/PrinceNorin/bakanovels/controllers/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var APIRouter *gin.Engine

func init() {
	am := buildAuthMiddleware()

	r := router.Get()
	api := r.Group("/api")
	api.Use(cors.Default())

	api.POST("/login", am.LoginHandler)

	r1 := api.Group("/v1")
	r1.Use(am.MiddlewareFunc())

	APIRouter = r
}
