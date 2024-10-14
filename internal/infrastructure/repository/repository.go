package repository

import "github.com/jackc/pgx/v5/pgxpool"

type DataRepository struct {
	db *pgxpool.Pool
}

func NewDataRepository(db *pgxpool.Pool) *DataRepository {
	return &DataRepository{db: db}
}

func (r DataRepository) Create() {
	r.db.Query()
}