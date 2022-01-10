package main

import (
	"fmt"
	"log"
	"net/http"
	"wireExp"
)

func main() {
	server:=wireExp.NewServer()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 9090),
		Handler: server.Router.GetRouter(),
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println()
	}
}