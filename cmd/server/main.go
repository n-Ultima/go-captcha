package main

import (
	"log"

	transportHTTP "github.com/PolyLmao/go-captcha/internal/transport/http"
)

// The struct for our app.
type App struct{}

// Setup method - Spin everything up.
func (app *App) Setup() {
	var handler transportHTTP.Handler
	log.Println("Setting up application on :8080")
	if err := handler.SetupRoutes(); err != nil {
		log.Fatalf("Could not setup routes: %v", err)
	}
}

// Server entry-point.
func main() {
	var app App
	app.Setup()
}
