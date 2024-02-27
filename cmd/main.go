package main

import (
	conf "currency-checker/config"
	api "currency-checker/internal/api"
	model "currency-checker/internal/database"
	currency "currency-checker/internal/services"

	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db = conf.ConnectDB()

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Cross{}, &model.Request{})

	go func() {
		router := gin.Default()
		router.POST("/updates", api.CreateRequestHandler)
		router.GET("/updates/:id", api.GetRequestHandler)
		// router.GET("/updates/", api.GetLatestUpdatedRequestHandler)
		router.Run("localhost:8082")
	}()

	go func() {
		ticker := time.NewTicker(5 * time.Second) // timeout -> config
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				currency.UpdateRequests(db)
			}
		}
	}()
	select {}
}
