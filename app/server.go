package app

import "github.com/gin-gonic/gin"

func router() *gin.Engine {
	r := gin.Default()
	r.GET("/messages/search", searchHandler)
	return r
}
