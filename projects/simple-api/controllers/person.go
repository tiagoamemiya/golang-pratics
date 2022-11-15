package controllers

import (
	"api/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func PersonIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.PersonsAll())
}

func PersonGet(w http.ResponseWriter, r *http.Request) {
	id := getRouteId(r)
	person := models.PersonGet(id)
	if person.Id == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(person)
}

func PersonNew(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)
	models.PersonNew(&person)

	json.NewEncoder(w).Encode(person)
}

func PersonDelete(w http.ResponseWriter, r *http.Request) {
	id := getRouteId(r)
	p := models.PersonGet(id)
	if p.Id == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	models.PersonDelete(p, id)
	w.WriteHeader(http.StatusNoContent)
}

func PersonUpdate(w http.ResponseWriter, r *http.Request) {
	id := getRouteId(r)
	p := models.PersonGet(id)
	if p.Id == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&p)
	models.PersonUpdate(&p)

	json.NewEncoder(w).Encode(&p)
}

func getRouteId(r *http.Request) int {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	return id
}
