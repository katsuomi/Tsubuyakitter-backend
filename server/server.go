package server

import (
	"github.com/gin-gonic/gin"
	"github.com/katsuomi/gin-like-twitter-api/controllers"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
		u.GET("/:id", ctrl.Show)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	p := r.Group("/posts")
	{
		ctrl := controllers.PostController{}
		p.GET("", ctrl.Index)
		p.POST("", ctrl.Create)
		p.GET("/:id", ctrl.Show)
		p.PUT("/:id", ctrl.Update)
		p.DELETE("/:id", ctrl.Delete)
	}

	return r
}
