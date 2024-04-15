package main

import (
	"loja-alura/routes"
	"net/http"
)

func main() {
	routes.CarregaTodasAsRotas()
	http.ListenAndServe(":8000", nil)
}
