package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app := mux.NewRouter()

	dep := InitDependencies()

	SetupRoutes(app, dep)

	fmt.Println("Server listening at :8080 ")
	return http.ListenAndServe(":8080", app)
}
