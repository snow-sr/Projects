package main

import (
	"github.com/snow-sr/learningGo/database"
	"github.com/snow-sr/learningGo/server"
)

func main() {
	database.StartDb()
	server := server.NewServer()
	server.Run()
}
