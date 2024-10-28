package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)


type CrawlerRepository struct {
	WordList        *WordList
	URLList         *URLList
	WordLocation    *WordLocation
	LinkBetweenURLs *LinkBetweenURLs
	LinkWord        *LinkWord
}

func NewCrawlerRepository(db *pgxpool.Pool) *CrawlerRepository {
	return &CrawlerRepository{
		WordList:        NewWordList(db),
		URLList:         NewURLList(db),
		WordLocation:    NewWordLocation(db),
		LinkBetweenURLs: NewLinkBetweenURLs(db),
		LinkWord:        NewLinkWord(db),
	}
}
