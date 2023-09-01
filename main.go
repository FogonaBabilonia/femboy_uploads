package main

import (
	"github.com/FogonaBabilonia/femboy_uploads/config"
	"github.com/FogonaBabilonia/femboy_uploads/database"
	"github.com/FogonaBabilonia/femboy_uploads/routes"
	"github.com/joho/godotenv"
)

func load_env() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func init() {
	load_env()
	config.Load_key()
	database.ConnectDatabase()
}

func main() {
	server := routes.NewServer(":3000")
	server.Run()
}
