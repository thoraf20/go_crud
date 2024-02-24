package routes

import (
	controller "golang-restaurant-management/controller"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menuId", controller.GetMenu())
	incomingRoutes.PATCH("/menus/:menuId", controller.UpdateMenu())
}