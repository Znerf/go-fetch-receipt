package controllers

import (
	// "fmt"
	"receipt/api/services"
	"receipt/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"time"
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
				"error": "No receipt found for that ID.",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "The receipt is invalid.",
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
	
	var validate *validator.Validate
	validate = validator.New()
    validate.RegisterValidation("date", validateDateFormat)
	validate.RegisterValidation("time", validateTimeFormat)


	err := validate.Struct(receipt)
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	receipt, err = cc.ReceiptService.Create(receipt)
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

//Need to add somewhere else
func validateDateFormat(fl validator.FieldLevel) bool {
    dateStr := fl.Field().String()
    _, err := time.Parse("2006-01-02", dateStr)
    return err == nil
}

func validateTimeFormat(fl validator.FieldLevel) bool {
    dateStr := fl.Field().String()
    _, err := time.Parse("15:04", dateStr)
    return err == nil
}