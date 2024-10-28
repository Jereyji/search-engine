package repository

import (
	"context"

	"github.com/Jereyji/search-engine/internal/domain/entity"
)

type WordList interface {
	Create(context context.Context, wordList *entity.WordList) (int, error)
	Word(context context.Context, id int) (*entity.WordList, error)
	Update(context context.Context, wordList *entity.WordList) error
	Delete(context context.Context, id int) error
}

type URLList interface {
	Create(context context.Context, url *entity.URLList) (int, error)
	URL(context context.Context, link string) (*entity.URLList, error)
	Update(context context.Context, url *entity.URLList) error
	Delete(context context.Context, id int) error
}

type WordLocation interface {
	Create(context context.Context, wordLocation *entity.WordLocation) (int, error)
	WordLocation(context context.Context, id int) (*entity.WordLocation, error)
	Update(context context.Context, wordLocation *entity.WordLocation) error
	Delete(context context.Context, id int) error
}

type LinkBetweenURLs interface {
	Create(context context.Context, linkBetweenURL *entity.LinkBetweenURL) (int, error)
	LinkBetweenURLs(context context.Context, id int) (*entity.LinkBetweenURL, error)
	Update(context context.Context, linkBetweenURL *entity.LinkBetweenURL) error
	Delete(context context.Context, id int) error
}

type LinkWord interface {
	Create(context context.Context, linkWord *entity.LinkWord) (int, error)
	LinkWord(context context.Context, id int) (*entity.LinkWord, error)
	Update(context context.Context, linkWord *entity.LinkWord) error
	Delete(context context.Context, id int) error
}

type CrawlerRepository struct {
	WordList
	URLList
	WordLocation
	LinkBetweenURLs
	LinkWord
}
