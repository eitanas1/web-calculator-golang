package application

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/bulbosaur/web-calculator-golang/pkg/calc"
)

type Config struct {
	Addr string
}

func GettingConfig() *Config {
	config := new(Config)
	address := os.Getenv("PORT")
	config.Addr = address
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response := ErrorResponse{
			Error: "Internal server error",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	result, err := calc.Calc(request.Expression)
	if err != nil {
		response := ErrorResponse{
			Error: "Expression is not valid",
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response)
		return
	}
	responce := Response{
		Result: result,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responce)
}

func RunServer() {
	address := ":" + GettingConfig().Addr
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Internal server error")
	}
}
