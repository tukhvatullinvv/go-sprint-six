package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func Run(logger *log.Logger) (*http.Server, error) {
	if logger == nil {
		return nil, fmt.Errorf("empty logger")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handlers.HandlerUpload)
	mux.HandleFunc("/", handlers.HandlerMain)
	serv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return serv, nil
}
