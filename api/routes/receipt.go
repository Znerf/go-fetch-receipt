package routes

import (
	"receipt/api/controllers"
	"receipt/resource"
)

type ReceiptRoutes struct {
	router           resource.Router
	receiptController controllers.ReceiptController
}

func NewReceiptRoutes(
	router resource.Router,
	receiptController controllers.ReceiptController,

) ReceiptRoutes {
	return ReceiptRoutes{
		router:           router,
		receiptController: receiptController,
	}
}

func (c ReceiptRoutes) Setup() {
	
	receiptGroup := c.router.Gin.Group("/receipts")
	{
		receiptGroup.GET("/", c.receiptController.GetReceipt) // for testing 
		receiptGroup.POST("/process", c.receiptController.CreateReceipt)
		receiptGroup.GET("/:id/points", c.receiptController.GetPoints)
	}
}