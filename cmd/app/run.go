package app

import (
	config "auth/config"
	http2 "auth/controllers/http"
	"auth/controllers/http/middleware"
	"auth/infrastructure/postgres"
	"auth/internal/repositories"
	"auth/internal/usecases"
	"auth/pkg/jwt"
	pkg "auth/pkg/services"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

var (
	cfg            *config.Config
	postgresClient *postgres.Client
	redisClient    *redis.Client

	sessionService pkg.SessionService
	hashService    pkg.HashService

	accountRepository repositories.AccountRepository
	sessionRepository repositories.SessionRepository

	signUpUseCase usecases.SignUpUseCase
	signInUseCase usecases.SignInUseCase
)

func Run() {
	var err error
	cfg, err = config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	initPostgres()
	initServices()
	initRepositories()
	initUseCases()

	defer redisClient.Close()
	defer postgresClient.Close()

	runServer()
}

func initServices() {
	hashService = pkg.NewHashService()

	accessTokenService := jwt.NewTokenService(cfg.JWT.SignSecretToken)
	refreshTokenService := jwt.NewTokenService(cfg.JWT.SignSecretToken)
	sessionService = pkg.NewSessionService(cfg.TokenConfig, accessTokenService, refreshTokenService)
}

func runServer() {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	mw := middleware.NewMiddleware()

	http2.NewSignUpController(router, signUpUseCase, mw)
	http2.NewSignInController(router, signInUseCase, mw)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	fmt.Println("postgres migrate success")
}

func initRepositories() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	accountRepository = CreatePGAccountRepository(postgresClient)
	sessionRepository = CreateSessionRepository(redisClient)
}

func initUseCases() {
	signUpUseCase = usecases.NewSignUpUseCase(
		accountRepository,
		sessionRepository,
		sessionService,
		hashService,
	)
	signInUseCase = usecases.NewSignInUseCase(
		accountRepository,
		sessionRepository,
		sessionService,
		hashService,
	)

}
