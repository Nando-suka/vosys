package main

import (
	"log"
	"voting-system/database"
	"voting-system/routes"
)

func main() {
	// Connect to database
	db := database.ConnectDatabase()

	// Mengatur routing
	r := routes.SetupRouter(db)

	// Running server
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
