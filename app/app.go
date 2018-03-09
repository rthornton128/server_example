package app

import (
	"net/http"
	"github.com/rthornton128/test/handlers"
)

// App is the application as a whole
type App struct{}

// NewApp create and returns a new application
func NewApp() *App {
	return &App{}
}

// RegisterHandlers for the application
func (a *App) RegisterHandlers() {
	http.Handle("/user/", handlers.AccessLogger(handlers.NewUserHandler()))
}
