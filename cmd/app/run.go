package app

import (
	config "auth/config"
	http2 "auth/controllers/http"
	"auth/infrastructure/postgres"
	"auth/internal/usecases"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	cfg            *config.Config
	postgresClient *postgres.Client

	signUpUseCase usecases.SignUpUseCase
)

func Run() {
	var err error
	cfg, err = config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	initPostgres()
	initUseCases()

	defer postgresClient.Close()

	runServer()
}

func runServer() {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	http2.NewSignUpController(router, signUpUseCase)

	address := fmt.Sprintf("%s:%s", cfg.Http.Host, cfg.Http.Port)
	fmt.Printf("starting server at %s\n", address)

	err := http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
}

func initPostgres() {
	var err error

	fmt.Println("starting postgres client")
	postgresClient, err = postgres.NewClient(cfg.PG)
	if err != nil {
		fmt.Printf("postgres client error: %v", err)
		return
	}

	err = postgresClient.MigrateUp()
	if err != nil {
		if errors.Is(err, postgres.ErrNoChange) {
			fmt.Println("nothing to migrate")
			return
		}
		fmt.Printf("postgres migrate error: %v", err)
		return
	}
}

func initUseCases() {
	signUpUseCase = usecases.NewSignUpUseCase()
}
