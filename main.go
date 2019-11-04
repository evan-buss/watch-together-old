package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/evan-buss/watch-together/server"
	"github.com/evan-buss/watch-together/server/chat"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	s := &server.Server{Router: http.NewServeMux(), Hub: chat.NewHub()}
	go s.Hub.Run()
	s.Routes()
	if err := http.ListenAndServe(":8080", s.Router); err != nil {
		return errors.Wrap(err, "server listener")
	}
	return nil
}
