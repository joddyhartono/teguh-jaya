package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joddyhartono/teguh-jaya/internal/config"
	"github.com/joddyhartono/teguh-jaya/internal/database"
	"github.com/joddyhartono/teguh-jaya/internal/handler"
	"github.com/joddyhartono/teguh-jaya/internal/repository"
	"github.com/joddyhartono/teguh-jaya/internal/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("database connected")

	productRepo := repository.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepo)

	r := router.New(router.Handlers{
		Product: productHandler,
	})

	srv := &http.Server{
		Addr:			":8080",
		Handler: 		r,
		ReadTimeout: 	10 * time.Second,
		WriteTimeout: 	30 * time.Second,
		IdleTimeout:	time.Minute,
	}
	
	log.Printf("server running on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}