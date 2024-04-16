package controllers

import (
	"log"
	"loja-alura/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		strNome := r.FormValue("nome")
		strDescricao := r.FormValue("descricao")
		strPreco := r.FormValue("preco")
		strQuantidade := r.FormValue("quantidade")

		floatPreco, err := strconv.ParseFloat(strPreco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		intQuantidade, err := strconv.Atoi(strQuantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriaNovoProduto(strNome, strDescricao, floatPreco, intQuantidade)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		strId := r.FormValue("id")
		strNome := r.FormValue("nome")
		strDescricao := r.FormValue("descricao")
		strPreco := r.FormValue("preco")
		strQuantidade := r.FormValue("quantidade")

		intId, err := strconv.Atoi(strId)
		if err != nil {
			log.Println("Erro na conversão do Id para int:", err)
		}

		floatPreco, err := strconv.ParseFloat(strPreco, 64)
		if err != nil {
			log.Println("Erro na conversão do Preço para float64:", err)
		}

		intQuantidade, err := strconv.Atoi(strQuantidade)
		if err != nil {
			log.Println("Erro na conversão da Quantidade para int:", err)
		}

		models.AtualizaProduto(intId, strNome, strDescricao, floatPreco, intQuantidade)
	}
	http.Redirect(w, r, "/", 301)
}
