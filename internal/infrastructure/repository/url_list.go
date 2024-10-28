package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/Jereyji/search-engine/internal/domain/entity"
	"github.com/Jereyji/search-engine/internal/infrastructure/repository/queries"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLList struct {
	db *pgxpool.Pool
	mu sync.RWMutex
}

func NewURLList(db *pgxpool.Pool) *URLList {
	return &URLList{db: db}
}

func (s *URLList) Create(context context.Context, url *entity.URLList) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddUrlList, url.Link, url.Is_parsed).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *URLList) URL(context context.Context, link string) (*entity.URLList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var url entity.URLList
	err := s.db.QueryRow(context, queries.GetUrlList, link).Scan(&url.ID, &url.Link)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &url, nil
		}

		return nil, err
	}

	return &url, nil
}

func (s *URLList) Update(context context.Context, url *entity.URLList) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.UpdateUrlList, url.Link, url.Is_parsed, url.ID)
	return err
}

func (s *URLList) Delete(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteUrlList, id)
	return err
}
