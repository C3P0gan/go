package routes

import (
    "net/http"
    "loja-alura/controllers"
)

func CarregaTodasAsRotas() {
	http.HandleFunc("/", controllers.Index)
    http.HandleFunc("/new", controllers.New)
    http.HandleFunc("/insert", controllers.Insert)
}
