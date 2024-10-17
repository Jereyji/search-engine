package repository_interface

import (
	"context"

	"github.com/Jereyji/search-engine/internal/domain/entity"
)

type SearchRepository interface {
	AddWordList(context context.Context, wordList *entity.WordList) (int, error)
	DeleteWord(context context.Context, id int) error
	GetWord(context context.Context, id int) (*entity.WordList, error)

	AddURL(context context.Context, url *entity.URLList) (int, error)
	DeleteURL(context context.Context, id int) error
	GetURL(context context.Context, id int) (*entity.URLList, error)

	AddWordLocation(context context.Context, wordLocation *entity.WordLocation) (int, error)
	DeleteWordLocation(context context.Context, id int) error
	GetWordLocation(context context.Context, id int) (*entity.WordLocation, error)

	AddLinkBetweenURLs(context context.Context, linkBetweenURL *entity.LinkBetweenURL) (int, error)
	DeleteLinkBetweenURLs(context context.Context, id int) error
	GetLinkBetweenURLs(context context.Context, id int) (*entity.LinkBetweenURL, error)

	AddLinkWord(context context.Context, linkWord *entity.LinkWord) (int, error)
	DeleteLinkWord(context context.Context, id int) error
	GetLinkWord(context context.Context, id int) (*entity.LinkWord, error)
}
