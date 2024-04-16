package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
    models.Personalidades = []models.Personalidade {
        {Nome: "Nome 1", Historia: "Historia 1"},
        {Nome: "Nome 2", Historia: "Historia 2"},
    }

	fmt.Println("Iniciando o servidor REST com Go")
	routes.HandleRequest()
}
