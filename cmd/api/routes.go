package main

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(app *mux.Router, dep *dependencies) {
	app.HandleFunc("/health", dep.configServerController.Health).Methods("GET")
}
