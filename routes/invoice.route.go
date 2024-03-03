package routes

import (
	controller "golang-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoiceId", controller.GetInvoice())
	incomingRoutes.PATCH("/invoices/:invoiceId", controller.UpdateInvoice())
}