package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"blondshuffler/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Get("/foo", handlers.HandlerFoo)

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)

	http.ListenAndServe(listenAddr, router)
	fmt.Println("hello world")
}
