package main

import cfgctrl "go-web/internal/app/config_server/controller"

type dependencies struct {
	configServerController *cfgctrl.ConfigServerController
}

func InitDependencies() *dependencies {

	configServerCtrl := cfgctrl.NewConfigServiceController()

	return &dependencies{
		configServerController: configServerCtrl,
	}
}
