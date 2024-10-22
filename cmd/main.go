package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/Jereyji/search-engine/internal/application/service"
	"github.com/Jereyji/search-engine/internal/infrastructure/repository"
	"github.com/Jereyji/search-engine/internal/pkg/config"
	router "github.com/Jereyji/search-engine/internal/pkg/console_router"
	"github.com/Jereyji/search-engine/internal/pkg/postgres"
	"github.com/Jereyji/search-engine/internal/pkg/server"
	"github.com/Jereyji/search-engine/internal/presentation/handler"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ConfigDataBase struct {
	Port     string `env:"DB_PORT" env-default:"5432"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
	User     string `env:"DB_USERNAME" env-default:"user"`
	Password string `env:"DB_PASSWORD"`
	SSLMode  string `env:"DB_SSLMODE" env-default:"disable"`
}

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err = godotenv.Load("deployments/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfgDataBase ConfigDataBase
	if err := cleanenv.ReadEnv(&cfgDataBase); err != nil {
		log.Fatal("Error reading environment variables: ", err)
	}

	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfgDataBase.User, cfgDataBase.Password, cfgDataBase.Host, cfgDataBase.Port, cfgDataBase.Name, cfgDataBase.SSLMode,
	)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	postgresDB, err := postgres.NewPostgresDB(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}
	defer postgresDB.Close()

	repos := repository.NewCrawlerRepository(postgresDB)
	service := service.NewCrawlerService(repos)
	handler := handler.NewCrawlerHandler(service)

	router := router.NewRouter()
	initializeRoutes(router, handler)

	fmt.Println("Server is running...")
	var key contextKey
	key = "key"
	server.ListenAndServe(context.WithValue(ctx, key, cfg.DataLinks), router)
}

type contextKey string

const (
	crawlCommand = "crawl"
	getLinkText  = "get-link-text"
	getListLinks = "get-list-links"
)

func initializeRoutes(router *router.Router, handler *handler.CrawlerHandler) {
	router.HandleFunc(crawlCommand, handler.Crawl)
}
