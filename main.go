package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/keyslapperdev/task-manager-mono/server/router"
	"github.com/keyslapperdev/task-manager-mono/server/storage"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r := router.SetupRouter(storage.NewDBStorer(true))

	r.Run(":" + port)
}
