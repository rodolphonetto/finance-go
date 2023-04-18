package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
	"testes.com/packages/repositories"
	"testes.com/packages/types"
	"testes.com/packages/utils"
	"testes.com/packages/validation"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var data types.Category

	if err := utils.ValidateJSONDecode(w, r, &data); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(data)
	validationErrors, result := validation.ValidateData(&data)

	if !result {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErrors)
		return
	}

	var id *gorm.DB
	var err error

	id, err = repositories.CreateCategory(data.Name)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Houve um erro ao criar categoria: %v", err)
	}

	fmt.Fprintf(w, "Received JSON data: %v", id.Value)
}
