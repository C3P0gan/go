package main

import (
	"api-gin-rest/database"
	"api-gin-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
