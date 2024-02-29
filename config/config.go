package conf

import (
	model "currency-updater/internal/database"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	path := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable Timezone=UTC",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.", err)
	}

	log.Println("Successfully connected to database.")
	db.AutoMigrate(&model.Cross{}, &model.Request{})

	return db
}

// func ConnectDB() *gorm.DB {
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return db
// }
