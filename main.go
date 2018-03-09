package main

import (
	"log"
	"net/http"
	"github.com/rthornton128/test/app"
)

func main() {
	app := app.NewApp()
	app.RegisterHandlers()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}