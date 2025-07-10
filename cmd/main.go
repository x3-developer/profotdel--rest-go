package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os/signal"
	"profotdel-rest/config"
	"profotdel-rest/internal/container"
	"profotdel-rest/internal/router"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	loadEnv()
	cfg := config.LoadConfig()
	diContainer := container.NewContainer(cfg, ctx, &wg)

	srv := newHTTPServer(cfg, diContainer)
	runServer(srv)

	<-ctx.Done()
	gracefulShutdown(srv, &wg)
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file")
	}
}

func newHTTPServer(cfg *config.Config, diContainer *container.Container) *http.Server {
	r := router.NewRouter(cfg, diContainer)

	return &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}
}

func runServer(srv *http.Server) {
	go func() {
		log.Printf("starting server on :%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("stopped listening server: %v", err)
		}
	}()
}

func gracefulShutdown(srv *http.Server, wg *sync.WaitGroup) {
	log.Println("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("error shutting down server: %v", err)
	}

	log.Println("waiting for background goroutines to finish...")
	wg.Wait()
	log.Println("server gracefully stopped")
}
