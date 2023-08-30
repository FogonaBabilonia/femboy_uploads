package main

import (
	"github.com/FogonaBabilonia/femboy_uploads/database"
	"github.com/FogonaBabilonia/femboy_uploads/routes"
)

func init() {
	database.ConnectDatabase()
}

func main() {
	server := routes.NewServer(":3000")
	server.Run()
}
