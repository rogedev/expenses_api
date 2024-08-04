package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rogedev/expenses_api/service/auth"
	"github.com/rogedev/expenses_api/types"
	"github.com/rogedev/expenses_api/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleRegister).Methods("POST")
}

func (h *Handler) HandleLogin(responseWriter http.ResponseWriter, request *http.Request) {}

func (h *Handler) HandleRegister(responseWriter http.ResponseWriter, request *http.Request) {

	var payload RegisterUserPayload

	if err := utils.ParseJSON(request, &payload); err != nil {
		utils.WriteError(responseWriter, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(responseWriter, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)

	if err == nil {
		utils.WriteError(
			responseWriter,
			http.StatusBadRequest,
			fmt.Errorf("user with email %s already exists", payload.Email),
		)

		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	if err != nil {
		utils.WriteError(responseWriter, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(responseWriter, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(responseWriter, http.StatusCreated, nil)
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=125"`
}
