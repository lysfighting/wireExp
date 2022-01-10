package router

import (
	"github.com/gin-gonic/gin"
	"wireExp/internal/api/handler"
)

type Router struct{
	router *gin.Engine
}

func NewRouter(userHandler *handler.UserHandler) *Router {
	var router *gin.Engine
	gin.SetMode(gin.DebugMode)
	router = gin.Default()

	router.GET("/", userHandler.Index)
	return &Router{router: router}
}

func (r *Router) GetRouter()  *gin.Engine{
	return r.router
}