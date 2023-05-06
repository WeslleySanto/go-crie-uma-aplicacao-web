package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func openConncetion() *sql.DB {
	db, err := sql.Open("sqlite3", "./alura_loja.db")
	checkErr(err)
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func getProducts() []Produto {
	db := openConncetion()

	getAllProducts, err := db.Query("select * from produtos")
	checkErr(err)

	p := Produto{}
	produtos := []Produto{}

	for getAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := getAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		checkErr(err)

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	// produtos := []Produto{
	// 	{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
	// 	{"Tenis", "Confortável", 89, 3},
	// 	{"Fone", "Muito bom", 59, 2},
	// 	{"Produto novo", "Muito legal", 1.99, 1},
	// }

	produtos := getProducts()

	tmpl.ExecuteTemplate(w, "Index", produtos)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
