package application

import (
	"log"
	"net/http"

	"github.com/bulbosaur/web-calculator-golang/internal/config"
	"github.com/bulbosaur/web-calculator-golang/internal/handlers"
)

func RunServer() {
	config := config.GettingConfig()
	address := ":" + config.Addr
	http.HandleFunc("/api/v1/calculate", handlers.CalcHandler)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Internal server error")
	}
}
