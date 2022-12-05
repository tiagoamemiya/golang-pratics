package main

import (
	"api/routes"
	"log"
)

func main() {
	log.Println("Starting server...")
	routes.HandleRequests()
}
