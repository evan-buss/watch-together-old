package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/evan-buss/watch-together/server"
	"github.com/evan-buss/watch-together/server/chat"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	s := &server.Server{Router: chi.NewMux(), Hub: chat.NewHub()}
	go s.Hub.Run()
	s.Routes()

	// Make sure connections don't take too long
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      s.Router,
	}

	if err := server.ListenAndServe(); err != nil {
		return errors.Wrap(err, "server listener")
	}
	return nil
}
