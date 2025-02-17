package main

import (
	"github.com/prolgrammer/BM_authService/cmd/app"
	_ "github.com/prolgrammer/BM_authService/docs"
)

// @title Auth Service
// @version 1.0
// @description service for auth users

// @host localhost:8080
func main() {
	app.Run()
}
