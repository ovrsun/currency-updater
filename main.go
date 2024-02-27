package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// use env
var key = "bf8ab72c635dc35b3d27a5da"
var url_pair = "https://v6.exchangerate-api.com/v6/" + key + "/pair/"

type Cross struct {
	ID   string `json:"id"`
	Code string `json:"code"` // "EUR/MXN"
}

type Request struct {
	ID      int       `json:"id"`
	Code    string    `json:"code"`
	Updated time.Time `json:"updated"`
	Rate    float64   `json:"rate"`
}

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Cross{}, &Request{})

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
				updateRequests()
			}
		}
	}()
	select {}
}

func FindCurrency(code string) (Cross, error) {
	var cross Cross

	if result := db.Where("Code = ?", code).First(&cross); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Cross{}, fmt.Errorf("currency with code %s not found", code)
		}
		return Cross{}, result.Error
	}
	return cross, nil
}

func GetNotUpdatedRequests() []Request {
	var requests []Request
	db.Where("rate = ?", 0.0).Find(&requests)
	return requests
}

func createRequestHandler(ctx *gin.Context) {
	var request Request

	if err := ctx.BindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "cant create new request"}) // TODO remove str constants in code
		return
	}

	res, err := FindCurrency(request.Code)
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

	var request = Request{ID: id}

	if result := db.Find(&request); result.Error != nil { // TODO need checking gorm error not found, else internal error
		ctx.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &request)
}

func getLatestUpdatedRequestHandler(ctx *gin.Context) {
	var request Request
	code := ctx.Query("code")
	code2 := url.QueryEscape(code)
	db.Order("updated DESC").First(&request)
	ctx.JSON(http.StatusOK, &request)
}

func updateRequests() { // TODO mb make it parallels, init workers count?
	requests := GetNotUpdatedRequests()
	if len(requests) == 0 {
		log.Println("Nothing to update")
		return
	}

	for _, request := range requests {
		base, target := splitCodeIntoPair(request.Code) // eur/usd
		rate, err := getCurrencyRate(base, target)
		db.Model(&request).Select("updated", "rate").Updates(map[string]interface{}{"updated": time.Now().UTC(), "rate": rate})
	}
	log.Printf("Successfully updated %d row(s)", len(requests))
}

func splitCodeIntoPair(code string) (string, string) { // TODO validation? add check for code (e.g. there is no '/' or smth like that)
	codes := strings.Split(code, "/")
	base := codes[0]
	target := codes[1]
	return base, target
}

func getCurrencyRate(base, target string) (float64, error) {
	url := url_pair + strings.ToLower(base) + "/" + strings.ToLower(target) // TODO printf
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		// log.Println("error: ", err)
		return 0, err
	}

	// client := http.Client{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// log.Println("error: ", err)
		return 0, err
	}

	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body) // io
	// err = json.Unmarshal(body, &result)
	var result map[string]interface{} // TODO make struct with response structure

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	// if err != nil {
	// 	log.Println("error: ", err)
	// 	return 0.0
	// }

	return result["conversion_rate"].(float64), nil // struct field, nil
}
