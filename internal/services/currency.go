package currency

import (
	model "currency-updater/internal/database"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
)

var APIKey, APIurl string

func SetAPIKey() {
	APIKey = os.Getenv("API_KEY")
	APIurl = os.Getenv("API_URL")
}

func validateCode(code string) bool {
	if len(code) > 7 {
		return false
	}

	if contains_slash := strings.Contains(code, "/"); !contains_slash {
		return false
	}
	return true
}

func SplitCodeIntoPair(code string) (string, string, error) {
	if valid_code := validateCode(code); !valid_code {
		return "", "", errors.New("invalid currencies code")
	}

	codes := strings.Split(code, "/")
	base := codes[0]
	target := codes[1]
	return base, target, nil
}

func UpdateRequests(db *gorm.DB) {
	// TODO mb make it parallels, init workers count?
	requests := model.GetNotUpdatedRequests()
	if len(requests) == 0 {
		log.Println("Nothing to update")
		return
	}

	for _, request := range requests {
		base, target, _ := SplitCodeIntoPair(request.Code)
		rate, _ := getCurrencyRate(base, target)
		db.Model(&request).Select("updated", "rate").Updates(map[string]interface{}{"updated": time.Now().UTC(), "rate": rate})
	}
	log.Printf("Successfully updated %d row(s)", len(requests))
}

func getCurrencyRate(base, target string) (float64, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(APIurl, APIKey, strings.ToLower(base), strings.ToLower(target)),
		nil,
	)

	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var result map[string]interface{} // TODO make struct with response structure

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result["conversion_rate"].(float64), nil
}
