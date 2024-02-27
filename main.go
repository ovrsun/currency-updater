package main

import (
	conf "currency-checker/config"
	model "currency-checker/internal/database"
	currency "currency-checker/internal/services"

	"log"
	"net/http"
	"strconv"
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
		router.POST("/updates", createRequestHandler)
		router.GET("/updates/:id", getRequestHandler)
		router.GET("/updates/", getLatestUpdatedRequestHandler)
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

func createRequestHandler(ctx *gin.Context) {
	var request model.Request

	if err := ctx.BindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "cant create new request"}) // TODO remove str constants in code
		return
	}

	res, err := model.FindCurrency(request.Code)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Fffffffuuuuu!!11"})
		return
	}
	log.Println(res) // TODO: try 2 use slog or zap logger

	if result := db.Create(&request); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()}) // TODO dont return DB errors, log it
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"request id": &request.ID})
}

func getRequestHandler(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id")) // TODO validation

	var request = model.Request{ID: id}

	if result := db.Find(&request); result.Error != nil { // TODO need checking gorm error not found, else internal error
		ctx.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &request)
}

func getLatestUpdatedRequestHandler(ctx *gin.Context) {
	var request model.Request
	db.Order("updated DESC").First(&request)
	ctx.JSON(http.StatusOK, &request)
}
