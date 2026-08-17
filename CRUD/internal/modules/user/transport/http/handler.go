package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"GO-Crud/internal/modules/user/application"
	"GO-Crud/internal/modules/user/domain"
)

type Handler struct {
	service *application.UserService
}

func NewHandler(service *application.UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.service.Create(r.Context(), user); err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, user)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "failed to get users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var req UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user := &domain.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.service.Update(r.Context(), user); err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
