package main

import (
	"github.com/erica7dev/api-rest-golang/database"
	"github.com/erica7dev/api-rest-golang/server"
)

func main() {
	database.StartDB()
	
	server := server.NewServer()

	server.Run()
}