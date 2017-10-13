package controllers

import (
	"github.com/PrinceNorin/bakanovels/controllers/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
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
	{
		r1.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"id":    1,
				"name":  "Norin",
				"email": "norin@example.com",
			})
		})
	}

	APIRouter = r
}
