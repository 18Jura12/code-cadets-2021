package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const betByIdPath = "/bet/:id"
const betsByUserIdPath = "/user/:id/bets"
const betsByStatusPath = "/bets?status={parameter}"

type WebServer struct {
	router             *gin.Engine
	port               int
	readWriteTimeoutMs int
}

func NewServer(port, readWriteTimeoutMs int, controller Controller) *WebServer {
	server := &WebServer{
		router:             gin.Default(),
		port:               port,
		readWriteTimeoutMs: readWriteTimeoutMs,
	}
	server.registerRoutes(controller)
	return server
}

func (w *WebServer) Start(ctx context.Context) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", w.port),
		Handler:      w.router,
		ReadTimeout:  time.Duration(w.readWriteTimeoutMs) * time.Millisecond,
		WriteTimeout: time.Duration(w.readWriteTimeoutMs) * time.Millisecond,
	}
	errs := make(chan error)

	go func() {
		err := server.ListenAndServe()
		errs <- err
	}()

	log.Printf("Started http server, port: %s, host: %s\n", w.port, "127.0.0.1")

	select {
	case err := <-errs:
		log.Printf("An error occurred: %s", err.Error())
		return

	case <-ctx.Done():
		ctx, clear := context.WithTimeout(context.Background(), 1*time.Second)
		defer clear()

		// gracefully shutdown server
		err := server.Shutdown(ctx)

		if err != nil {
			log.Printf("An error occurred: %s", err.Error())
		}
		return
	}
}

func (w *WebServer) registerRoutes(controller Controller) {
	w.router.GET(betByIdPath, controller.FetchBetById())
	w.router.GET(betsByUserIdPath, controller.FetchBetsByUserId())
	w.router.GET(betsByStatusPath, controller.FetchBetsByStatus())
}

// Controller handles api calls
type Controller interface {
	FetchBetById() gin.HandlerFunc
	FetchBetsByUserId() gin.HandlerFunc
	FetchBetsByStatus() gin.HandlerFunc
}