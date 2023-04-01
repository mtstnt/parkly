package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"parkly/config"

	"github.com/gin-gonic/gin"
)

func startApp() error {
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)

	router := gin.Default()

	server := &http.Server{
		Addr:        fmt.Sprintf("%s:%d", config.AppHost, config.AppPort),
		Handler:     router,
		ReadTimeout: 30 * time.Second,
		ErrorLog:    logger,
	}

	if err := setupDependencies(); err != nil {
		return err
	}

	if err := setupRoutes(router); err != nil {
		return err
	}

	signalReceiver := make(chan os.Signal, 1)
	errReceiver := make(chan error, 1)
	signal.Notify(signalReceiver, os.Interrupt, syscall.SIGTERM)

	go func() {
		errReceiver <- server.ListenAndServe()
	}()
	logger.Println("Server started")

	select {
	case sig := <-signalReceiver:
		logger.Println("Signal " + sig.String() + " is received. Terminating program gracefully.")
	case err := <-errReceiver:
		logger.Println("Error occurred: " + err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Exited successfully")
	return nil
}

func main() {
	log.Fatalln(startApp())
}
