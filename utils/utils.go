package utils

import (
	"encoding/json"
	"net/http"
)

func ValidateJSONDecode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}
