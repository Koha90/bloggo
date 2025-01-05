package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/koha90/bloggo/internal/storage"
	"github.com/koha90/bloggo/internal/types"
)

func (s *ApiServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	req := new(types.CreateUserRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return fmt.Errorf("invalid loaded data (JSON)")
	}

	if req.Username == "" || req.Password == "" {
		return fmt.Errorf("username and password are required")
	}

	user, err := types.NewUser(req.Username, req.FirstName, req.LastName, req.Password)
	if err != nil {
		return fmt.Errorf("invalid user data")
	}

	err = s.store.CreateUser(user)
	if errors.Is(err, storage.ErrUserAlreadyExists) {
		return fmt.Errorf("user with this username already exists")
	}
	if err != nil {
		return fmt.Errorf("could not create user")
	}

	return WriteJSON(w, http.StatusCreated, user)
}

func (s *ApiServer) handleUserByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	user, err := s.store.UserByID(uint(id))
	if err != nil {
		return fmt.Errorf("user not found")
	}
	return WriteJSON(w, http.StatusOK, user)
}

func (s *ApiServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	var req types.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	user, err := s.store.UserByUsername(req.Username)
	if err != nil {
		return fmt.Errorf("no username??: %w", err)
	}

	if !user.ValidatePassword(req.Password) {
		return fmt.Errorf("not authenticated")
	}

	token, err := createJWT(user)
	if err != nil {
		return err
	}

	resp := types.LoginResponse{
		Token:    token,
		Username: user.Username,
	}

	return WriteJSON(w, http.StatusOK, resp)
}

// utilitars
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	// временный буфер для кодирования.
	buf, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	_, err = w.Write(buf)

	return err
}

func createJWT(user *types.User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"username":  user.Username,
	}

	secret := "ajkl;sjdf"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

//
// type HTTPError struct {
// 	StatusCode int    `json:"-"`
// 	Message    string `json:"message"`
// }
//
// func (e *HTTPError) Error() string {
// 	return e.Message
// }

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}

func getID(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid ID")
	}

	return id, nil
}
