package services

import (
	"fmt"
	"receipt/resource"
	"receipt/models"
	"strconv"
	"unicode"
	"math"
	"time"
	"strings"
)

type ReceiptService struct {
	db     resource.Database
}

func NewReceiptService(db resource.Database) ReceiptService {
	return ReceiptService{
		db:     db,
	}
}

func (c ReceiptService) Create(receipt models.Receipt) (models.Receipt, error) {
	var total float64 =0
	for _, x := range receipt.Items {
		num, err :=  strconv.ParseFloat(x.Price, 64) //strconv.Atoi(x.Price)
		if err != nil {
			fmt.Println("Error:", err)
			return receipt, err
		}
		
		total = total + num
        fmt.Println(x)
    }
	receipt.Total = fmt.Sprintf("%.2f", total )
	return receipt, c.db.DB.Create(&receipt).Error
}

func (c ReceiptService) Get() ([]models.Receipt, error) {
	var receipts []models.Receipt
	err := c.db.DB.Model(&models.Receipt{}).Find(&receipts).Error
	return receipts,err
}

func (c ReceiptService) countAlphaNumeric(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func (c ReceiptService) GetPoints( ID models.BINARY16) (int, error) {

	receipt := models.Receipt{}
	pointCount := 0
	
	err := c.db.DB.Where("id = ?", ID).Preload("Items").First(&receipt).Error
	if err != nil{
		return pointCount,err
	}

	pointCount = c.countAlphaNumeric(receipt.Retailer)

	num, _:= strconv.ParseFloat(receipt.Total, 64) 
	decimal := int((num - math.Floor(num))*100)

	if decimal == 0 {
		pointCount = pointCount +50
	}
	if decimal % 25 ==0 {
		pointCount = pointCount + 25
	}
	pointCount = pointCount + len(receipt.Items) / 2 * 5
	for _, x := range receipt.Items {
		price, _ := strconv.ParseFloat(x.Price,64)
		if len(strings.TrimSpace(x.ShortDescription)) % 3 == 0 {
			pointCount = pointCount + int(math.Ceil(0.2 * price ))
		}
	}
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return pointCount, err
	}

	day := date.Day()
	if day%2 != 0 {
		pointCount = pointCount + 6
	} 

	parsedTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return pointCount, err
	}
	startTime, _ := time.Parse("15:04", "14:00")
	endTime, _ := time.Parse("15:04", "16:00")
	if parsedTime.After(startTime) && parsedTime.Before(endTime) {
		pointCount = pointCount + 10
	} 
	
	return pointCount, err

}
