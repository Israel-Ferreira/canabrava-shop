package main

import (
	"log"
	"net/http"

	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/internal/routes"
)

func main() {

	porta := ":8080"

	router := routes.InitRoutes()

	log.Println("Servidor iniciado na porta ", porta)
	log.Fatalln(http.ListenAndServe(porta, router))
}
