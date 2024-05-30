package handler

import (
	"net/http"

	"github.com/andrearcaina/echo/pkg/utils"
	"github.com/go-chi/chi/v5"
)

type CustomHandler struct{}

func (h *CustomHandler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/test", h.Test)

	return r
}

func (h *CustomHandler) Test(w http.ResponseWriter, r *http.Request) {
	utils.SendJSON(w, http.StatusOK, map[string]string{"message": "testing custom handler endpoint"})
}
