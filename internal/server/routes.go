package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/koha90/bloggo/internal/types"
)

func (s *ApiServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	req := new(types.CreateUserRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	user, err := types.NewUser(req.Username, req.FirstName, req.LastName, req.Password)
	if err != nil {
		return err
	}
	if err := s.store.CreateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, user)
}

func (s *ApiServer) handleUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		id, err := getID(r)
		if err != nil {
			return err
		}

		user, err := s.store.UserByID(uint(id))
		if err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, user)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

// utilitars
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func getID(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}
