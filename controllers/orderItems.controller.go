package controller

import (
	"github.com/gin-gonic/gin"
)

func GetOrderItems() gin.HandleFunc() {
	return func(c *gin.Context) {

	}
}

func GetOrderItemsByOrder() gin.HandleFunc() {
	return func(c *gin.Context) {
		
	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	return func(c *gin.Context) {
		
	}
}

func GetOrderItem() gin.HandleFunc() {
	return func(c *gin.Context) {

	}
}

func CreateOrderItem() gin.HandleFunc() {
	return func(c *gin.Context) {

	}
}

func UpdateOrderItem() gin.HandleFunc() {
	return func(c *gin.Context) {
		
	}
}