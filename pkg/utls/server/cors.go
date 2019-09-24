package server

import (
	"github.com/gin-gonic/gin"
)

//PublicCors returns authorization for diferentes cors
func (s *Server) PublicCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Next()
	}
}

//PrivateCors is responsible only for routes that are pivate
func (s *Server) PrivateCors() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
