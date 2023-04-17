package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"testes.com/packages/types"
	"testes.com/packages/utils"
	"testes.com/packages/validation"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var data types.Pessoa

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

	fmt.Fprintf(w, "Received JSON data: %v", data)
}
