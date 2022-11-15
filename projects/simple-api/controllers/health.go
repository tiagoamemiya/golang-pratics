package controllers

import (
	"encoding/json"
	"net/http"
)

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	content := map[string]string{"message": "Alive!"}
	json.NewEncoder(w).Encode(content)
}
