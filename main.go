package main

import (
	"github.com/FogonaBabilonia/femboy_uploads/routes"
)

func main() {
	server := routes.NewServer(":3000")
	server.Run()
}
