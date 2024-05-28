package routes

import (
	"net/http"

	"github.com/andrearcaina/echo/pkg/utils"
)

func InitAPI() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		test := map[string]string{"message": "Hello, World!"}

		utils.SendJSON(w, http.StatusOK, test)
	})

	return mux
}
