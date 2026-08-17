package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(
	router *mux.Router,
	handler *Handler,
) {
	router.HandleFunc(
		"/users",
		handler.CreateUser,
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/users",
		handler.GetUsers,
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/users/{id}",
		handler.GetUser,
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/users/{id}",
		handler.UpdateUser,
	).Methods(http.MethodPut)

	router.HandleFunc(
		"/users/{id}",
		handler.DeleteUser,
	).Methods(http.MethodDelete)
}