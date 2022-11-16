package main

import "go-web/internal/app/config_server"

type dependencies struct {
	configServerController *config_server.ConfigServerController
}

func InitDependencies() *dependencies {

	configServerCtrl := config_server.NewConfigServiceController()

	return &dependencies{
		configServerController: configServerCtrl,
	}
}
