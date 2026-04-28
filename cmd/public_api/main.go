package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	publicapi "github.com/d-darac/lagra/internal/api/public_api/v1"
	"github.com/d-darac/lagra/internal/database"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Print("[main] Connecting to database...")
	db, err := database.New(ctx, database.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "lagra-test",
		Password: "mysecretpassword",
		DBName:   "lagra-test",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("[main] Error connecting to database: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	publicapi.RegisterHandlers(mux, db)

	api := http.NewServeMux()
	api.Handle("/v1/", http.StripPrefix("/v1", mux))

	server := &http.Server{
		Handler:           api,
		Addr:              ":8080",
		ReadHeaderTimeout: time.Second * 15,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}

	go func() {
		log.Print("[main] Server starting...")
		var err error
		err = server.ListenAndServe()

		if err != http.ErrServerClosed {
			log.Printf("[main] Server failed to start: %v.", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	fmt.Println()
	log.Print("[main] Server stopping...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Println()
		log.Printf("[main] Server forced to shutdown: %v.", err)
		return
	}

	log.Print("[main] Server stopped.")
}
