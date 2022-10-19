package httphelper

import (
	"fmt"
	"net/http"
)

func JsonResponse(statusCode int, w http.ResponseWriter, body interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, body)
}
