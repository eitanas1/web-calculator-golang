package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bulbosaur/web-calculator-golang/internal/models"
	"github.com/bulbosaur/web-calculator-golang/pkg/calc"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(models.Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response := models.ErrorResponse{
			Error: "Internal server error",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	result, err := calc.Calc(request.Expression)
	if err != nil {
		response := models.ErrorResponse{
			Error: "Expression is not valid",
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := models.Response{
		Result: result,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
