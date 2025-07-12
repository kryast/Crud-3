package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-3.git/handlers"
)

func SetupRouter(customerHandler *handlers.CustomerHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/customers", customerHandler.CreateCustomer)

	return r
}
