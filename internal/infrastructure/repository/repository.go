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

type CrawlerRepository struct {
	db *pgxpool.Pool
	mu sync.RWMutex
}

func NewCrawlerRepository(db *pgxpool.Pool) *CrawlerRepository {
	return &CrawlerRepository{db: db}
}

func (s *CrawlerRepository) AddWordList(context context.Context, wordList *entity.WordList) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddWordList, wordList.Word, wordList.IsFiltred).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CrawlerRepository) DeleteWord(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteWordList, id)
	return err
}

func (s *CrawlerRepository) GetWord(context context.Context, id int) (*entity.WordList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var wordList entity.WordList
	err := s.db.QueryRow(context, queries.GetWordList, id).Scan(&wordList)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &wordList, nil
}

func (s *CrawlerRepository) AddURL(context context.Context, url *entity.URLList) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddUrlList, url.Link).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CrawlerRepository) DeleteURL(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteUrlList, id)
	return err
}

func (s *CrawlerRepository) GetURL(context context.Context, link string) (*entity.URLList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var url entity.URLList
	err := s.db.QueryRow(context, queries.GetUrlList, link).Scan(&url)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}

func (s *CrawlerRepository) AddWordLocation(context context.Context, wordLocation *entity.WordLocation) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddWordLocation, wordLocation.WordID, wordLocation.URLID, wordLocation.Location).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CrawlerRepository) DeleteWordLocation(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteWordLocation, id)
	return err
}

func (s *CrawlerRepository) GetWordLocation(context context.Context, id int) (*entity.WordLocation, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var wordLocation entity.WordLocation
	err := s.db.QueryRow(context, queries.GetWordLocation, id).Scan(&wordLocation)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &wordLocation, nil
}

func (s *CrawlerRepository) AddLinkBetweenURLs(context context.Context, linkBetweenURL *entity.LinkBetweenURL) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddLinkBetweenURL, linkBetweenURL.FromURLID, linkBetweenURL.ToURLID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CrawlerRepository) DeleteLinkBetweenURLs(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteLinkBetweenURL, id)
	return err
}

func (s *CrawlerRepository) GetLinkBetweenURLs(context context.Context, id int) (*entity.LinkBetweenURL, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var linkBetweenURL entity.LinkBetweenURL
	err := s.db.QueryRow(context, queries.GetLinkBetweenURL, id).Scan(&linkBetweenURL)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &linkBetweenURL, nil
}

func (s *CrawlerRepository) AddLinkWord(context context.Context, linkWord *entity.LinkWord) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int
	err := s.db.QueryRow(context, queries.AddLinkWord, linkWord.WordID, linkWord.LinkID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CrawlerRepository) DeleteLinkWord(context context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(context, queries.DeleteLinkWord, id)
	return err
}

func (s *CrawlerRepository) GetLinkWord(context context.Context, id int) (*entity.LinkWord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var linkWord entity.LinkWord
	err := s.db.QueryRow(context, queries.GetLinkWord, id).Scan(&linkWord)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &linkWord, nil
}
