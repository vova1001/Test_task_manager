package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	r "github.com/vova1001/Test_task_manager/router"
)

func main() {
	logChan := make(chan string, 1000)
	go func() {
		for msg := range logChan {
			log.Println(msg)
		}
	}()
	mux := http.NewServeMux()
	mux.Handle("/tasks/", http.StripPrefix("/tasks", http.HandlerFunc(r.RegisterRouterTask(logChan))))
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Signal caught")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatalf("Server err shutdown: %v", err)
		}
		log.Println("Server stopped")
	}()
	log.Println("Server is running")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}

}
