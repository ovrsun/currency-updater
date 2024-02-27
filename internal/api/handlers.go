package api

import (
	model "currency-checker/internal/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRequestHandler(ctx *gin.Context) {
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

	// if result := db.Create(&request); result.Error != nil {

	if result := model.CreateRequest(&request); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()}) // TODO dont return DB errors, log it
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"request id": &request.ID})
}

func GetRequestHandler(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id")) // TODO validation

	var request = model.Request{ID: id}

	if result := model.FindRequest(&request); result.Error != nil { // TODO need checking gorm error not found, else internal error
		ctx.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &request)
}

// func getRequestHandler(ctx *gin.Context) {
// 	id, _ := strconv.Atoi(ctx.Param("id")) // TODO validation

// 	var request = model.Request{ID: id}

// 	if result := db.Find(&request); result.Error != nil { // TODO need checking gorm error not found, else internal error
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, &request)
// }

// func getLatestUpdatedRequestHandler(ctx *gin.Context) {
// 	var request model.Request
// 	db.Order("updated DESC").First(&request)
// 	ctx.JSON(http.StatusOK, &request)
// }