package api

import "github.com/gin-gonic/gin"

const betPath = "/bet"

type WebServer struct {
	router             *gin.Engine
	port               int
	readWriteTimeoutMs int
}




