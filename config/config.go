package conf

import (
	model "currency-updater/internal/database"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Timeout = 1 * time.Minute

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
