// Project structure:
//
// github-oauth-demo/
// ├── go.mod
// ├── go.sum
// ├── main.go
// ├── config/
// │   └── config.go
// ├── handlers/
// │   └── deploy.go
// ├── middleware/
// │   └── auth.go
// └── routes/
//     └── routes.go

// File: main.go
package main

import (
	"akila-i/github-oauth-go/config"
	"akila-i/github-oauth-go/routes"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	router := routes.SetupRoutes(cfg)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
