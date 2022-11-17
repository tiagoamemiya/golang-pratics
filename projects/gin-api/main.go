package main

import (
	"gin-api/config"
	"gin-api/routes"

	"github.com/gin-gonic/gin"
)

const port = ":8081"

func main() {
	config.DbConnect()
	r := gin.Default()
	routes.HandleRequest(r)
	r.Run(port)
}
