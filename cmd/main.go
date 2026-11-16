package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "APP: ", log.LstdFlags)
	serv, err := server.Run(logger)
	if err != nil {
		logger.Printf("server runnig error: %v", err)
		os.Exit(1)
	}
	if err1 := serv.ListenAndServe(); err1 != nil {
		logger.Printf("server runnig error: %v", err1)
		os.Exit(1)
	}
}
