package api

import (
	model "currency-updater/internal/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRequestHandler(ctx *gin.Context) {
	var request model.Request

	if err := ctx.BindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Can't create new request"}) // TODO remove str constants in code
		return
	}

	res, err := model.FindCurrenciesCode(request.Code)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "The currency does not exist"})
		return
	}
	log.Println(res) // TODO: try 2 use slog or zap logger

	if result := model.CreateRequest(&request); result.Error != nil {
		log.Printf("Failed to to create request: %v", result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
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

func getLatestUpdatedRequestHandler(ctx *gin.Context) {
	var request model.Request

	code := ctx.Request.URL.Query()["code"][0]
	model.GetLatestUpdateRequest(code, &request)

	ctx.JSON(http.StatusOK, &request)
}
