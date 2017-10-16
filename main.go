package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PrinceNorin/bakanovels/config"
	"github.com/PrinceNorin/bakanovels/controllers"
	"github.com/PrinceNorin/bakanovels/utils/locale"
)

func main() {
	c := config.Get()
	err := locale.InitLocale(c.I18n)
	if err != nil {
		log.Fatal(err.Error())
	}

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
