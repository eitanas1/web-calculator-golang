package server

import (
	"log"
	"net/http"

	"github.com/bulbosaur/web-calculator-golang/config"
	"github.com/bulbosaur/web-calculator-golang/internal/handlers"
)

func RunServer() {
	config := config.GettingConfig()
	address := ":" + config.Addr

	http.HandleFunc("/api/v1/calculate", handlers.CalcHandler)
	http.HandleFunc("/coffe", handlers.CoffeeHandler)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Internal server error")
	}
}
