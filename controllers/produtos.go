package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/WeslleySanto/go-crie-uma-aplicacao-web/models"
	produtoModels "github.com/WeslleySanto/go-crie-uma-aplicacao-web/models"
	utils "github.com/WeslleySanto/go-crie-uma-aplicacao-web/utils"
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
	if utils.IsPost(r) {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := utils.ConvertStringToFloat64(r.FormValue("preco"))
		quantidade := utils.ConvertStringToInt(r.FormValue("quantidade"))

		models.Create(nome, descricao, preco, quantidade)
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
	if utils.IsPost(r) {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := utils.ConvertStringToFloat64(r.FormValue("preco"))
		quantidade := utils.ConvertStringToInt(r.FormValue("quantidade"))

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		produtoModels.AtualizaProduto(idConvertidaParaInt, nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}
