package router

import (
	"myapp/controller"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	r.POST("/post/create", controller.PostCreate)
	r.PUT("/post/update", controller.PostUpdate)
	r.DELETE("/post/delete", controller.PostDelete)
	r.GET("/posts", controller.PostGetAll)
	r.GET("/post", controller.PostGetByID)
}
