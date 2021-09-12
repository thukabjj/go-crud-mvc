package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/thukabjj/go-crud-mvc/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome, descricao := r.FormValue("nome"), r.FormValue("descricao")
		preco, err := strconv.ParseFloat(strings.TrimSpace(r.FormValue("preco")), 64)
		validateError("Erro na conversão do preço", err)
		quantidade, err := strconv.Atoi(strings.TrimSpace(r.FormValue("quantidade")))
		validateError("Erro na conversão da qauntidade", err)

		produto := models.Produto{
			Nome:       nome,
			Descricao:  descricao,
			Preco:      preco,
			Quantidade: quantidade,
		}

		models.CriarNovoProduto(&produto)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto, err := strconv.Atoi(strings.TrimSpace(r.URL.Query().Get("id")))
	validateError("Erro na conversão do id do produto", err)
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto, err := strconv.Atoi(strings.TrimSpace(r.URL.Query().Get("id")))
	validateError("Erro na conversão do id do produto", err)
	produto := models.BuscarProdutoPorId(idProduto)
	templates.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nome, descricao := r.FormValue("nome"), r.FormValue("descricao")

		idProduto, err := strconv.Atoi(strings.TrimSpace(r.FormValue("id")))
		validateError("Erro na conversão do id do produto", err)

		quantidade, err := strconv.Atoi(strings.TrimSpace(r.FormValue("quantidade")))
		validateError("Erro na conversão do quantidade do produto", err)

		preco, err := strconv.ParseFloat(strings.TrimSpace(r.FormValue("preco")), 64)
		validateError("Erro na conversão do preço do produto", err)

		produto := models.Produto{
			Id:         idProduto,
			Nome:       nome,
			Descricao:  descricao,
			Preco:      preco,
			Quantidade: quantidade,
		}
		models.AtualizarProduto(&produto)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func validateError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
