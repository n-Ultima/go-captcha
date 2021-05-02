package http

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/PolyLmao/go-captcha/internal/models"
)

// Our captcha model.
var captcha models.Captcha

// Our base URL for testing.
var baseURL = "http://localhost:8080"

// Test our "/captcha" endpoint.
// Check the status-code to ensure the request went through.
// In the next test, we check if the captcha image was created along with the successful request.
func TestNewCaptcha(t *testing.T) {
	request, _ := http.Get(baseURL + "/captcha")
	if request.StatusCode != http.StatusOK {
		t.Errorf("Did not recieve expected response from /captcha: %v", request.StatusCode)
	}
	json.NewDecoder(request.Body).Decode(&captcha)
}

// Test our "/images" endpoint.
// Check the status-code to ensure the path is valid.
// If the path is invalid for any reason, we can assume that something isn't right.
func TestServeImage(t *testing.T) {
	request, _ := http.Get(baseURL + "/images/" + captcha.Code)
	if request.StatusCode != http.StatusOK {
		t.Errorf("Did not recieve expected response from /images/%s: %v", captcha.Code, request.StatusCode)
	}
}
