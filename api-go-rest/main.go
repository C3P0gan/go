package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor REST com Go")
	routes.HandleRequest()
}
