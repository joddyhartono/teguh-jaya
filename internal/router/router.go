package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joddyhartono/teguh-jaya/internal/handler"
)

type Handlers struct {
	Product *handler.ProductHandler
	Category *handler.CategoryHandler
}

func New(h Handlers) http.Handler {
	r := chi.NewRouter()
	
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 30))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:		[]string{"http://localhost:5173"},
		AllowedMethods:		[]string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:		[]string{"*"},
		AllowCredentials:	false,
	}))

	r.Use(middleware.StripSlashes)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			r.Get("/", h.Product.List)
			r.Get("/{id}", h.Product.GetByID)
		})
		r.Route("/categories", func(r chi.Router) {
			r.Get("/", h.Category.List)
		})
	})

	return r
}	