package main

import (
	"github.com/bulbosaur/web-calculator/internal/application"
)

func main() {
	app := application.New()
	app.RunServer
}
