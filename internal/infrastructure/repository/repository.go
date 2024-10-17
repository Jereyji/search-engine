package repository

import (
	"context"

	"github.com/Jereyji/search-engine/internal/domain/entity"
	"github.com/Jereyji/search-engine/internal/infrastructure/Repository/queries"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DataRepository struct {
	db *pgxpool.Pool
}

func NewDataRepository(db *pgxpool.Pool) *DataRepository {
	return &DataRepository{db: db}
}

func (s *DataRepository) AddWordList(context context.Context, wordList *entity.WordList) (int, error) {
	var id int
	err := s.db.QueryRow(context, queries.AddWordList, wordList.Word, wordList.IsFiltred).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *DataRepository) DeleteWord(context context.Context, id int) error {
	_, err := s.db.Exec(context, queries.DeleteWordList, id)
	return err
}

func (s *DataRepository) GetWord(context context.Context, id int) (*entity.WordList, error) {
	var wordList entity.WordList
	err := s.db.QueryRow(context, queries.GetWordList, id).Scan(&wordList)
	if err != nil {
		return nil, err
	}

	return &wordList, nil
}

// ...................

func (s *DataRepository) AddURL(context context.Context, url *entity.URLList) (int, error) {
	var id int
	err := s.db.QueryRow(context, queries.AddUrlList, url.Link).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *DataRepository) DeleteURL(context context.Context, id int) error {
	_, err := s.db.Exec(context, queries.DeleteUrlList, id)
	return err
}

func (s *DataRepository) GetURL(context context.Context, id int) (*entity.URLList, error) {
	var url entity.URLList
	err := s.db.QueryRow(context, queries.GetUrlList).Scan(&url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

// ...................

func (s *DataRepository) AddWordLocation(context context.Context, wordLocation *entity.WordLocation) (int, error) {
	var id int
	err := s.db.QueryRow(context, queries.AddWordLocation, wordLocation.WordID, wordLocation.URLID, wordLocation.Location).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *DataRepository) DeleteWordLocation(context context.Context, id int) error {
	_, err := s.db.Exec(context, queries.DeleteWordLocation, id)
	return err
}

func (s *DataRepository) GetWordLocation(context context.Context, id int) (*entity.WordLocation, error) {
	var wordLocation entity.WordLocation
	err := s.db.QueryRow(context, queries.GetWordLocation, id).Scan(&wordLocation)
	if err != nil {
		return nil, err
	}

	return &wordLocation, nil
}

// ...................

func (s *DataRepository) AddLinkBetweenURLs(context context.Context, linkBetweenURL *entity.LinkBetweenURL) (int, error) {
	var id int
	err := s.db.QueryRow(context, queries.AddLinkBetweenURL, linkBetweenURL.FromURLID, linkBetweenURL.ToURLID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *DataRepository) DeleteLinkBetweenURLs(context context.Context, id int) error {
	_, err := s.db.Exec(context, queries.DeleteLinkBetweenURL, id)
	return err
}

func (s *DataRepository) GetLinkBetweenURLs(context context.Context, id int) (*entity.LinkBetweenURL, error) {
	var linkBetweenURL entity.LinkBetweenURL
	err := s.db.QueryRow(context, queries.GetLinkBetweenURL, id).Scan(&linkBetweenURL)
	if err != nil {
		return nil, err
	}

	return &linkBetweenURL, nil
}

// ...................

func (s *DataRepository) AddLinkWord(context context.Context, linkWord *entity.LinkWord) (int, error) {
	var id int
	err := s.db.QueryRow(context, queries.AddLinkWord, linkWord.WordID, linkWord.LinkID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *DataRepository) DeleteLinkWord(context context.Context, id int) error {
	_, err := s.db.Exec(context, queries.DeleteLinkWord, id)
	return err
}

func (s *DataRepository) GetLinkWord(context context.Context, id int) (*entity.LinkWord, error) {
	var linkWord entity.LinkWord
	err := s.db.QueryRow(context, queries.GetLinkWord, id).Scan(&linkWord)
	if err != nil {
		return nil, err
	}

	return &linkWord, nil
}
