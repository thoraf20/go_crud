package routes

import (
	controller "golang-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItemId", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:orderId", controller.GetOrderItemsByOrder())
	incomingRoutes.PATCH("/orderItems/:orderItemId", controller.UpdateOrderItem())
}