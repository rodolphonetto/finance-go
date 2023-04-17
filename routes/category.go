package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"testes.com/packages/controllers"
)

func SetupCategoryRoutes(r *mux.Router) {
	r.HandleFunc("/create", controllers.CreateCategory).Methods(http.MethodPost)
}
