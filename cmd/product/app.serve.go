package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/lib/config"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/route"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// initiate config
	cfg, err := config.ConfigInit()
	if err != nil {
		panic(err)
	}

	// initiate database
	db, err := config.PostgresInit(cfg)
	if err != nil {
		panic(err)
	}

	// initiate gin
	app := gin.New()
	r := route.NewRoute(db, app, cfg)
	r.RouteInit()

	// create new http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: app.Handler(),
	}

	go func() {
		// service connections
		log.Println("Starting Server ...")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

}
