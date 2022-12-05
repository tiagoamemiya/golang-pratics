package routes

import (
	"api/controllers"
	"api/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const ROOT_PATH = "/api/v1"

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middlewares.SetJsonContentType)
	r.HandleFunc("/health", controllers.HomeIndex).Methods("Get")
	r.HandleFunc(ROOT_PATH+"/persons", controllers.PersonIndex).Methods("Get")
	r.HandleFunc(ROOT_PATH+"/persons", controllers.PersonNew).Methods("Post")
	r.HandleFunc(ROOT_PATH+"/persons/{id}", controllers.PersonGet).Methods("Get")
	r.HandleFunc(ROOT_PATH+"/persons/{id}", controllers.PersonDelete).Methods("Delete")
	r.HandleFunc(ROOT_PATH+"/persons/{id}", controllers.PersonUpdate).Methods("Put")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
