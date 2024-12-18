package handlers

import (
	"fmt"
	"net/http"
)

func CoffeeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	fmt.Fprintf(w, "I'm a teapot")
}
