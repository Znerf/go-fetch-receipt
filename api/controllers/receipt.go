package controllers

import (
	"fmt"
	"receipt/api/services"
	"receipt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReceiptController struct {
	ReceiptService services.ReceiptService
}

func NewReceiptController(
	ReceiptService services.ReceiptService,

) ReceiptController {
	return ReceiptController{
		ReceiptService: ReceiptService,
	}
}

//for viewing 
func (cc ReceiptController) GetReceipt(c *gin.Context) {
	receipt, err := cc.ReceiptService.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
        "data":  receipt,
    })
}

func (cc ReceiptController) GetPoints(c *gin.Context) {

	ID := c.Param("id")
	IdBinary, err := models.StringToBinary16(ID)

	points, err := cc.ReceiptService.GetPoints(IdBinary)
	if err != nil {
		if points == -1 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
        "points":  points,
    })
}


func (cc ReceiptController) CreateReceipt(c *gin.Context) {

	receipt := models.Receipt{}
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error() + "sdf",
		})
		return
	}
	fmt.Println("sdf", receipt)
	
	receipt, err := cc.ReceiptService.Create(receipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
        "id":  receipt.ID,
    })
}