package routes

import (
	controller "golang-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:orderId", controller.GetOrder())
	incomingRoutes.PATCH("/orders/:orderId", controller.UpdateOrder())
}