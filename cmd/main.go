package main

import (
	conf "currency-updater/config"
	api "currency-updater/internal/api"
	model "currency-updater/internal/database"
	init_db "currency-updater/internal/init"
	currency "currency-updater/internal/services"

	"time"

	"github.com/gin-gonic/gin"
)

var API_KEY, API_URL string

func main() {
	var db = conf.ConnectDB()
	model.SetDB(db)
	init_db.InitializeDB(db)
	currency.SetAPIKey()

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
