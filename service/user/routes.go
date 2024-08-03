package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleLogin).Methods("POST")
}

func (h *Handler) HandleLogin(responseWriter http.ResponseWriter, request *http.Request) {}

func (h *Handler) HandleRegister(responseWriter http.ResponseWriter, request *http.Request) {}
