package repository_interface

import (
	"context"

	"github.com/Jereyji/search-engine/internal/domain/entity"
)

type CrawlerRepositoryInterface interface {
	CreateWord(context context.Context, wordList *entity.WordList) (int, error)
	Word(context context.Context, id int) (*entity.WordList, error)
	UpdateWord(context context.Context, wordList *entity.WordList) error
	DeleteWord(context context.Context, id int) error
	
	CreateURL(context context.Context, url *entity.URLList) (int, error)
	URL(context context.Context, link string) (*entity.URLList, error)
	UpdateURL(context context.Context, url *entity.URLList) error
	DeleteURL(context context.Context, id int) error
	
	CreateWordLocation(context context.Context, wordLocation *entity.WordLocation) (int, error)
	WordLocation(context context.Context, id int) (*entity.WordLocation, error)
	UpdateWordLocation(context context.Context, wordLocation *entity.WordLocation) error
	DeleteWordLocation(context context.Context, id int) error
	
	CreateLinkBetweenURLs(context context.Context, linkBetweenURL *entity.LinkBetweenURL) (int, error)
	LinkBetweenURLs(context context.Context, id int) (*entity.LinkBetweenURL, error)
	UpdateLinkBetweenURLs(context context.Context, linkBetweenURL *entity.LinkBetweenURL) error
	DeleteLinkBetweenURLs(context context.Context, id int) error
	
	CreateLinkWord(context context.Context, linkWord *entity.LinkWord) (int, error)
	LinkWord(context context.Context, id int) (*entity.LinkWord, error)
	UpdateLinkWord(context context.Context, linkWord *entity.LinkWord) error
	DeleteLinkWord(context context.Context, id int) error

	PopularDomains(ctx context.Context) ([]entity.PopularDomain, error)
	PopularWords(ctx context.Context) ([]entity.PopularWord, error)
}
