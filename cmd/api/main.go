package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"os/signal"
	"profotdel-rest/config"
	"profotdel-rest/internal/shared/application/container"
	"profotdel-rest/internal/shared/application/router"
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
	diContainer := container.NewContainer(cfg, &wg)

	srv := newHTTPServer(cfg, diContainer)
	runServer(srv)

	<-ctx.Done()
	gracefulShutdown(srv, &wg)
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal("error loading .env file")
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
		logrus.Printf("starting server on :%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			logrus.Printf("stopped listening server: %v", err)
		}
	}()
}

func gracefulShutdown(srv *http.Server, wg *sync.WaitGroup) {
	logrus.Println("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("error shutting down server: %v", err)
	}

	logrus.Println("waiting for background goroutines to finish...")
	wg.Wait()
	logrus.Println("server gracefully stopped")
}
