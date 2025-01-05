// Package server - contains a server designer.
package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/koha90/bloggo/internal/storage"
)

// HTTPServer - server for work with http response and requests.
type ApiServer struct {
	port  string
	store *storage.Storage
}

// New - creating new HTTP server.
func NewHTTPServer(port string, store *storage.Storage) *http.Server {
	apiServer := &ApiServer{
		port:  port,
		store: store,
	}
	server := &http.Server{
		Addr:         port,
		Handler:      apiServer.RegisterRoutes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *ApiServer) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	s.setupAuthRoutes(r)
	s.setupUserRouter(r)

	return r
}

func (s *ApiServer) setupAuthRoutes(r chi.Router) {
	r.Post("/api/v1/signup", makeHTTPHandleFunc(s.handleCreateUser))
	r.Post("/api/v1/signin", makeHTTPHandleFunc(s.handleLogin))
}

func (s *ApiServer) setupUserRouter(r chi.Router) {
	r.Get("/api/v1/users/{id}", makeHTTPHandleFunc(s.handleUserByID))
}
