package routers

import (
	controller "transaction-stori/src/controller"

	"github.com/gin-gonic/gin"
)

// NewServer generate new server
func NewServer() *gin.Engine {

	server := gin.New()

	server.Use(gin.Logger())

	server.Use(gin.Recovery())
	return server
}

// RegisterRoutes register routes
func RegisterRoutes(server *gin.Engine) {

	server.GET("/health", controller.Health)

	server.POST("/transaction/send-email", controller.SendTranstactionsToEmail)

}
