package main

import (
	"log"
	"net/http"

	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/internal/config"
	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/internal/routes"
)

func init() {
	config.LoadEnv()
}

func main() {

	config.InitDb()

	porta := ":8080"

	router := routes.InitRoutes()

	log.Println("Servidor iniciado na porta ", porta)
	log.Fatalln(http.ListenAndServe(porta, router))
}
