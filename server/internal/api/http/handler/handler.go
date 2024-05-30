package http

import (
	"github.com/andrearcaina/echo/pkg/utils"
	"net/http"
)

type CustomHandler struct{}

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", CustomHandler{}.Test)

	return mux
}

func (h *CustomHandler) Test(w http.ResponseWriter, r *http.Request) {
	utils.SendJSON(w, http.StatusOK, map[string]string{"message": "testing API endpoint"})
}
