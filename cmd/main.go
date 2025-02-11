package main

import (
	"auth/cmd/app"
	_ "auth/docs"
)

// @title Auth Service
// @version 1.0
// @description service for auth users

// @host localhost:8080
func main() {
	app.Run()
}
