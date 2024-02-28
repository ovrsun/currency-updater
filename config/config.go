package conf

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// sqlDB, err := db.DB()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer sqlDB.Close()

	return db
}
