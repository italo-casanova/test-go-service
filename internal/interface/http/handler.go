package http

import (
	"encoding/json"
	"net/http"

	"hello-service/internal/application"
)

type Handler struct {
	helloUsecase *application.HelloUsecase
}

func NewHandler(uc *application.HelloUsecase) *Handler {
	return &Handler{helloUsecase: uc}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	msg := h.helloUsecase.GetHelloMessage()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
