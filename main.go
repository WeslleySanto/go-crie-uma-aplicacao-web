package main

import (
	"net/http"

	"github.com/WeslleySanto/go-crie-uma-aplicacao-web/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
