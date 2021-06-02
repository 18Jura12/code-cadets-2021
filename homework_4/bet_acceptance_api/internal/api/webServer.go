package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const betPath = "/bet"

type WebServer struct {
	router             *gin.Engine
	port               int
	readWriteTimeoutMs int
}

func NewServer(port, readWriteTimeoutMs int, ctrl Controller) *WebServer {
	server := &WebServer{
		router:             gin.Default(),
		port:               port,
		readWriteTimeoutMs: readWriteTimeoutMs,
	}
	server.registerRoutes(ctrl)
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

func (w *WebServer) registerRoutes(ctrl Controller) {
	w.router.POST(betPath, ctrl.CreateBet())
}

type Controller interface {
	CreateBet() gin.HandlerFunc
}
