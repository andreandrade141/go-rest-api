package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Client struct {
	id   int
	Name string
}

var Clients = []Client{}

func GetClients(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Clients)
}

func main() {
	router := gin.Default()
	router.GET("/Clients", GetClients)

	router.Run("localhost:8080")
}
