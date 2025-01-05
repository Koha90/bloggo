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
	NewApiServer := &ApiServer{
		port:  port,
		store: store,
	}
	server := &http.Server{
		Addr:         port,
		Handler:      NewApiServer.RegisterRoutes(),
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
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Post("/api/v1/signup", makeHTTPHandleFunc(s.handleCreateUser))
	r.Get("/api/v1/users/{id}", makeHTTPHandleFunc(s.handleUserByID))

	return r
}
