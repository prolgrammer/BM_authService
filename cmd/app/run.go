package app

import (
	config "auth/config"
	http2 "auth/controllers/http"
	"auth/controllers/http/middleware"
	"auth/infrastructure/postgres"
	"auth/internal/repositories"
	"auth/internal/usecases"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	cfg            *config.Config
	postgresClient *postgres.Client

	accountRepository repositories.AccountRepository

	signUpUseCase usecases.SignUpUseCase
	signInUseCase usecases.SignInUseCase
)

func Run() {
	var err error
	cfg, err = config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	//gin.SetMode(gin.DebugMode)

	initPostgres()
	initRepositories()
	initUseCases()

	defer postgresClient.Close()

	runServer()
}

func runServer() {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	mw := middleware.NewMiddleware()

	http2.NewSignUpController(router, signUpUseCase, mw)
	http2.NewSignInController(router, signInUseCase, mw)

	address := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	fmt.Printf("starting server at %s\n", address)

	fmt.Println("Current mode:", gin.Mode())
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

func initRepositories() {
	accountRepository = CreatePGAccountRepository(postgresClient)
}

func initUseCases() {
	signUpUseCase = usecases.NewSignUpUseCase(
		accountRepository,
	)
	signInUseCase = usecases.NewSignInUseCase(
		accountRepository,
	)

}
