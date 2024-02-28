package model

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFindCurrency(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	SetDB(db)

	if err != nil {
		t.Fatalf("Error while connecting to db: %s", err)
	}

	err = db.AutoMigrate(&Cross{})
	if err != nil {
		t.Fatalf("Error while migrating: %s", err)
	}

	testData := []Cross{
		{ID: "1", Code: "EUR/USD"},
		{ID: "2", Code: "USD/EUR"},
	}

	for _, data := range testData {
		if err = db.Create(&data).Error; err != nil {
			t.Fatalf("Error while creating test data: %s", err)
		}
	}

	code := "EUR/USD"
	cross, err := FindCurrency(code)
	if err != nil {
		t.Fatalf("Error while finding currency by code: %s", err)
	}

	expected := "EUR/USD"
	if cross.Code != expected {
		t.Errorf("Expected result: %s, got: %s", expected, cross.Code)
	}
}
