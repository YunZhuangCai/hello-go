package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
var Router *gin.Engine
func main(){
	Router = gin.New()
	Router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	Router.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"hello:":"world!"})
	})

	Router.Run("127.0.0.1:8888")
}
