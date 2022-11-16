package config_server

import (
	"go-web/internal/pkg/http_helper"
	"net/http"
)

type ConfigServerController struct{}

func NewConfigServiceController() *ConfigServerController { return &ConfigServerController{} }

func (ctrl *ConfigServerController) Health(w http.ResponseWriter, r *http.Request) {
	http_helper.JsonResponse(http.StatusOK, w, nil)
}
