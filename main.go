package main

import (
	"os"

	"github.com/FogonaBabilonia/femboy_uploads/config"
	"github.com/FogonaBabilonia/femboy_uploads/database"
	"github.com/FogonaBabilonia/femboy_uploads/routes"
	"github.com/joho/godotenv"
)

func load_env() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	config.Secret_key = os.Getenv("JWT_SECRET_KEY")
}

func init() {
	load_env()
	database.ConnectDatabase()
}

func main() {
	server := routes.NewServer(":3000")
	server.Run()
}
