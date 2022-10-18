package main

import (
	"fmt"
	cfgctrl "go-web/internal/app/config_server/controller"
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

type dependencies struct {
	configServerController *cfgctrl.ConfigServerController
}

func InitDependencies() *dependencies {

	configServerCtrl := cfgctrl.NewConfigServiceController()

	return &dependencies{
		configServerController: configServerCtrl,
	}
}
