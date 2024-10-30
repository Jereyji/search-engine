package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type CrawlerRepository struct {
	db *pgxpool.Pool
	mu sync.RWMutex
}

func NewCrawlerRepository(db *pgxpool.Pool) *CrawlerRepository {
	return &CrawlerRepository{
		db: db,
	}
}
