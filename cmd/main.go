package main

import (
	// "context"
	"log"
	// "os/signal"
	// "syscall"

	// "github.com/Jereyji/search-engine.git/internal/infrastructure/repository"
	"github.com/Jereyji/search-engine.git/internal/domain/service"
	"github.com/Jereyji/search-engine.git/internal/pkg/config"
	// "github.com/Jereyji/search-engine.git/internal/pkg/postgres"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	// defer cancel()

	// psqlDB, err := postgres.NewPostgresDB(ctx, cfg.DataBaseURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer psqlDB.Close()

  service.Crawl(cfg.DataLinks, 1)
	// repos := repository.NewDataRepository(psqlDB)

}
