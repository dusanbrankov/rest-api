package main

import (
	// "context"
	"log"
	// "net/http"
	// "os"
	// "os/signal"
	// "syscall"
	// "time"

	"github.com/dusanbrankov/rest-api/cmd/api"
	"github.com/dusanbrankov/rest-api/config"
	"github.com/dusanbrankov/rest-api/db"
)

func main() {
	cfg := config.App

	// Initialize database  connection and SQL queries instance
	db := db.Initialize()
	cfg.Infolog.Printf("connected to database '%s'\n", cfg.DatabaseConfig.Name)

	// Initialize server with the database connection and run it
	srv := api.NewAPIServer(":8080", db)
	cfg.Infolog.Printf("running server on port :8080 (%s)", cfg.Enivronment)
	log.Fatal(srv.Run())
}

// func runServer(db *sql.DB, cfg *config.AppConfig) {
// 	const port = ":8080"
//
// 	srv := api.NewAPIServer(port, db)
// 	if err := srv.Run(); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	srv := &http.Server{
// 		Addr:         port,
// 		Handler:      api.Routes(),
// 		ReadTimeout:  5*time.Second,
// 		WriteTimeout: 10*time.Second,
// 	}
//
// 	cfg.Infolog.Printf("running server on port %s (%s)", port, cfg.Enivronment)
// 	log.Fatal(srv.ListenAndServe())
//
// 	// Start the server in a new goroutine
// 	go func() {
// 		cfg.Infolog.Printf("running server on port %s (%s)", port, cfg.Enivronment)
// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("listen: %s\n", err)
// 		}
// 	}()
//
// 	// Handle graceful shutdown
// 	gracefulShutdown(srv, cfg.Infolog)
// }
//
// func gracefulShutdown(srv *http.Server, logger *log.Logger) {
// 	quit := make(chan os.Signal, 1)
// 	// Listen for SIGINT (Ctrl+C) and SIGTERM signals
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//
// 	// Block until a signal is received
// 	<-quit
// 	logger.Println("shutting down server...")
//
// 	// Create a context with a timeout to allow the server to shut down gracefully
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()
//
// 	// Shutdown the server, allowing in-progress requests to complete
// 	if err := srv.Shutdown(ctx); err != nil {
// 		logger.Fatalf("server forced to shutdown: %v", err)
// 	}
//
// 	logger.Println("server exiting")
// }
//

