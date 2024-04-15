package main

import (
	"net/http"
	"text/template"
    "database/sql"

    _ "github.com/lib/pq"  // dependência usada durante o tempo de execução do app
)

type Produto struct {
    Id int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func conectaComBancoDeDados() *sql.DB {
    conexao := "user=postgres dbname=alura_loja password=1234 host=localhost sslmode=disable"
    db, err := sql.Open("postgres", conexao)
    if err != nil {
        panic(err.Error())
    }
    return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
    db := conectaComBancoDeDados()

    selectDeTodosOdProdutos, err := db.Query(
        "select * from produtos ;",
    )
    if err != nil {
        panic(err.Error())
    }

    p := Produto{}
    produtos := []Produto{}

    for selectDeTodosOdProdutos.Next() {
        var id, quantidade int
        var nome, descricao string
        var preco float64

        err = selectDeTodosOdProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
        if err != nil {
            panic(err.Error())
        }

        p.Nome = nome
        p.Descricao = descricao
        p.Preco = preco
        p.Quantidade = quantidade

        produtos = append(produtos, p)

    }

	temp.ExecuteTemplate(w, "Index", produtos)
    defer db.Close()
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
