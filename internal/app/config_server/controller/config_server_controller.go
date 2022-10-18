package controller

import (
	"fmt"
	"net/http"
)

type ConfigServerController struct{}

func NewConfigServiceController() *ConfigServerController { return &ConfigServerController{} }

func (ctrl *ConfigServerController) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, map[string]interface{}{"status": "alive"})
}
