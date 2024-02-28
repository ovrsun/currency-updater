package model

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	ConnectionFailedError = "Failed to connect to db: %s"
	MigrationFailedError  = "Failed to migrate db: %s"
	CreationTestDataError = "Failed to create test data: %s"
)

func TestFindCurrenciesCode(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	SetDB(db)

	if err != nil {
		t.Fatalf(ConnectionFailedError, err)
	}

	err = db.AutoMigrate(&Cross{})
	if err != nil {
		t.Fatalf(MigrationFailedError, err)
	}

	testData := []Cross{
		{ID: "1", Code: "EUR/USD"},
		{ID: "2", Code: "USD/EUR"},
	}

	for _, data := range testData {
		if err = db.Create(&data).Error; err != nil {
			t.Fatalf(CreationTestDataError, err)
		}
	}

	code := "EUR/USD"
	cross, err := FindCurrenciesCode(code)
	if err != nil {
		t.Fatalf("Failed to find currency by code: %s", err)
	}

	expected := "EUR/USD"
	if cross.Code != expected {
		t.Errorf("Expected result: %s, got: %s", expected, cross.Code)
	}
}

func TestCreateRequest(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	SetDB(db)

	if err != nil {
		t.Fatalf(ConnectionFailedError, err)
	}

	err = db.AutoMigrate(&Cross{})
	if err != nil {
		t.Fatalf(MigrationFailedError, err)
	}
}
