package server

import (
	"github.com/gin-gonic/gin"
)

//Server Instantiates server
type Server struct {
	App *gin.Engine
}

//NewServer is responsible to return an server engine
func NewServer() *Server {
	r := gin.Default()
	app := &Server{r}
	return app
}
