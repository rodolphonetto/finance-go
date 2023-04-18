package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"testes.com/packages/routes"
	"testes.com/packages/storage"
)

var DB *gorm.DB

func startServer(router *mux.Router) {
	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	routes.SetupRoutes(r)
	storage.InitStorage()
	startServer(r)
}
