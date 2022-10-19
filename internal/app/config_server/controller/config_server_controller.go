package controller

import (
	httphelper "go-web/internal/pkg/http_helper"
	"net/http"
)

type ConfigServerController struct{}

func NewConfigServiceController() *ConfigServerController { return &ConfigServerController{} }

func (ctrl *ConfigServerController) Health(w http.ResponseWriter, r *http.Request) {
	httphelper.JsonResponse(http.StatusOK, w, nil)
}
