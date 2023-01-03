package main

import (
	"transaction-stori/src/routers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	server := routers.NewServer()

	routers.RegisterRoutes(server)

	server.Run(":8080")
}
