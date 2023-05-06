package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/WeslleySanto/go-crie-uma-aplicacao-web/models"
	produtoModels "github.com/WeslleySanto/go-crie-uma-aplicacao-web/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := produtoModels.GetAllProducts()

	tmpl.ExecuteTemplate(w, "Index", produtos)
}

func Create(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoToFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeToInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		models.Create(nome, descricao, precoToFloat, quantidadeToInt)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeleteById(id)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	tmpl.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na convesão do preço para float64:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na convesão da quantidade para int:", err)
		}

		produtoModels.AtualizaProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
