package main

import (
	conf "currency-checker/config"
	"currency-checker/internal/api"
	model "currency-checker/internal/database"
	currency "currency-checker/internal/services"

	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var db = conf.ConnectDB()
	model.SetDB(db)
	db.AutoMigrate(&model.Cross{}, &model.Request{})

	go func() {
		router := gin.Default()
		api.RegisterRotes(router)
		router.Run(":8081")
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
