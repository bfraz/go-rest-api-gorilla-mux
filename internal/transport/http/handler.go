package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Server  *http.Server
	Router  *mux.Router
	Service CommentService
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)
	h.Router.Use(TimeoutMiddleware)

	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})

	h.Router.HandleFunc("/api/v1/comment", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/v1/comments", h.GetComments).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.DeleteComment).Methods("DELETE")
	/*
		h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World")
		})
	*/
}

func (h *Handler) Serve() error {
	// wrap ListenAndServe in a go routine making it non-blocking operation
	// allowing us to continue execution of Serve
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	// create a channel of type os.Signal and size 1
	c := make(chan os.Signal, 1)
	// notify the channel for any interrupt signals
	signal.Notify(c, os.Interrupt)
	// block until signal is received
	<-c

	// create a context with a timeout of 15 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer the cancel function to be called after the function returns
	defer cancel()
	// shutdown the server with the created context
	h.Server.Shutdown(ctx)

	log.Println("graceful shutdown")
	return nil
}
