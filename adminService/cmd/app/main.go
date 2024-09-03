package main

import (
	"adminService/internal/configs"
	"adminService/internal/db"
	"adminService/internal/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlers.MainHandler)

	cfg := configs.Get()

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      m,
		WriteTimeout: cfg.Timeout * time.Second,
		ReadTimeout:  cfg.Timeout * time.Second,
	}

	db.Connect()
	db.Migrate()

	log.Println("Listening " + cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
