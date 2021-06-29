package main

import (
	"github.com/reafreitas1/booksAPI_Golang/database"
	"github.com/reafreitas1/booksAPI_Golang/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
