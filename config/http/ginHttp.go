package http

import (
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	router *gin.RouterGroup
)

func init() {
	server = gin.Default()
	router = server.Group("/conductor/v1/api/")

}

func GetServer() *gin.Engine {
	return server
}

func GetRouter() *gin.RouterGroup {
	return router
}
