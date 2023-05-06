package routes

import (
	"net/http"

	produtoControllers "github.com/WeslleySanto/go-crie-uma-aplicacao-web/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", produtoControllers.Index)
	http.HandleFunc("/create", produtoControllers.Create)
	http.HandleFunc("/insert", produtoControllers.Insert)
}