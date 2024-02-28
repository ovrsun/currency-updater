package api

import "github.com/gin-gonic/gin"

func RegisterRotes(router *gin.Engine) {
	router.POST("/updates", CreateRequestHandler)
	router.GET("/updates/:id", GetRequestHandler)
	router.GET("/updates/", getLatestUpdatedRequestHandler)
}
