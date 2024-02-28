package init_db

import (
	model "currency-updater/internal/database"

	"gorm.io/gorm"
)

func InitializeDB(db *gorm.DB) {
	var count int64
	db.Model(&model.Cross{}).Count(&count)

	if count == 0 {
		crosses := []model.Cross{
			{ID: "1", Code: "EUR/USD"},
			{ID: "2", Code: "EUR/MXN"},
			{ID: "3", Code: "EUR/RUB"},
			{ID: "4", Code: "USD/EUR"},
			{ID: "5", Code: "USD/MXN"},
			{ID: "6", Code: "USD/RUB"},
			{ID: "7", Code: "MXN/EUR"},
			{ID: "8", Code: "MXN/USD"},
			{ID: "9", Code: "MXN/RUB"},
			{ID: "10", Code: "RUB/EUR"},
			{ID: "11", Code: "RUB/USD"},
			{ID: "12", Code: "RUB/MXN"},
		}

		db.Create(&crosses)
	}
}
