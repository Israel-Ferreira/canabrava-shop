package routes

import (
	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/internal/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/"))

	router.Post("/api/v1/adegas", controllers.CadastrarAdega)

	return router
}
