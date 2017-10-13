package main

import (
	"fmt"
	"github.com/PrinceNorin/bakanovels/config"
	"github.com/PrinceNorin/bakanovels/controllers"
	"log"
	"net/http"
	"time"
)

func main() {
	c := config.Get()
	APIRouter := controllers.APIRouter

	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)

	s := &http.Server{
		Addr:         addr,
		Handler:      APIRouter,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Listening on: %s", addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
