// Package classification Captcha
//
// Documentation for Captcha API
//
//  Schemes: http
//  Version: 1.0.0
//  License: MIT http://opensource.org/licenses/MIT
//
//  Produces:
//  - application/json
//  - image/png
//
// swagger:meta
package http

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

// The struct for our handler - contains a pointer to a router.
type Handler struct {
	router *mux.Router
}

// Set up all of the routes for our application.
func (h *Handler) SetupRoutes() error {
	options := middleware.RedocOpts{SpecURL: "/swagger"}
	documentationMiddleware := middleware.Redoc(options, nil)
	h.router = mux.NewRouter()
	h.router.HandleFunc("/captcha", h.NewCaptcha).Methods("GET")
	h.router.HandleFunc("/images/{code}", h.ServeImage).Methods("GET")
	h.router.Handle("/docs", documentationMiddleware)
	h.router.HandleFunc("/swagger", h.ServeDocs)
	if err := http.ListenAndServe(":8080", h.router); err != nil {
		return err
	}
	return nil
}
