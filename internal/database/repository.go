package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func FindCurrency(code string) (Cross, error) {
	var cross Cross

	if result := db.Where("Code = ?", code).First(&cross); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Cross{}, fmt.Errorf("currency with code %s not found", code)
		}
		return Cross{}, result.Error
	}
	return cross, nil
}

func GetNotUpdatedRequests() []Request {
	var requests []Request
	db.Where("rate = ?", 0.0).Find(&requests)
	return requests
}

func CreateRequest(request *Request) *gorm.DB {
	return db.Create(&request)
}

func FindRequest(request *Request) *gorm.DB {
	return db.Find(&request)
}