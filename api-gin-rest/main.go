package main

import (
	"api-gin-rest/models"
	"api-gin-rest/routes"
)

func main() {
    models.Alunos = []models.Aluno {
        {Nome: "John Doe", CPF: "00000000000", RG: "0000000"},
        {Nome: "Jane Doe", CPF: "00000000000", RG: "0000000"},
    }
	routes.HandleRequests()
}
