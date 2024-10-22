package service

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Jereyji/search-engine/internal/domain/entity"
	"github.com/Jereyji/search-engine/internal/domain/repository_interface"
	"github.com/PuerkitoBio/goquery"
)

/*
FUNCS:
- addIndex(self, soup, url) - Индексирование одной страницы
- getTextOnly(self, text) - Получение текста страницы
- separateWords(self, text) - Разбиение текста на слова
- isIndexed(self, url) - Проиндексирован ли URL
- addLinkRef(self, urlFrom, urlTo, linkText) - Добавление ссылки с одной страницы на другую
- crawl(self, urlList, maxDepth=1) - сбора данных
- getEntryId(self, tableName, fieldName, value) - получение идентификатора и добавление записи (РАЗДЕЛИТЬ)
*/

type CrawlerService struct {
	repository repository_interface.CrawlerRepository
}

func NewCrawlerService(repository repository_interface.CrawlerRepository) *CrawlerService {
	return &CrawlerService{
		repository: repository,
	}
}

type DataLinks struct {
	Url      string `yaml:"url"`
	Selector string `yaml:"selector"`
	Text     string `yaml:"text"`
}

// Загрузка страницы из списка обхода;
// Проверка, есть ли такая страница уже в индексе, если страницы в индексе нет, то она добавляется;
// Выделение новые ссылки со страницы и добавляются в список обхода;
// Далее паука переходит к следующему документу.

func (s *CrawlerService) Crawl(context context.Context, data DataLinks, maxDepth int) error {
		res, err := http.Get(data.Url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return err
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return err
		}

		_, err = s.repository.GetURL(context, data.Url) // CheckURL
		if err != nil {
			return err
		}

		curURL := entity.URLList{
			Link: data.Url,
		}

		curURL.ID, err = s.repository.AddURL(context, &curURL) // Add to urlList(url)
		if err != nil {
			return err
		}

		var crawlErr error

		doc.Find(data.Selector).Each(func(i int, selection *goquery.Selection) {
			title := selection.Find(data.Text).Text()
			words := separateText(&title)
			fmt.Println(words)
			if words != nil {
				if crawlErr = s.indexingAndRecordingWords(context, words, curURL.ID); crawlErr != nil {
					return
				}
			}

			link, exists := selection.Attr("href") 
			if exists {
				// if link[0] == '/' {
				// 	link = 
				// }

				_, err = s.repository.GetURL(context, link)
				if err != nil {
					crawlErr = err
					return
				}

				nextURL := entity.URLList {
					Link: link,
				}

				nextURL.ID, err = s.repository.AddURL(context, &entity.URLList{Link: nextURL.Link})
				if err != nil {
					crawlErr = err
					return
				}

				_, err = s.repository.AddLinkBetweenURLs(context, &entity.LinkBetweenURL{FromURLID: curURL.ID, ToURLID: nextURL.ID})
				if err != nil {
					crawlErr = err
					return
				}
			}
		})

		if crawlErr != nil {
			return crawlErr
		}
		// return
		// url
		// count words
		// count addedUrls

	return nil
}

func separateText(text *string) []string {
	if text == nil || *text == "" {
		return nil
	}

	words := strings.Fields(*text)
	return words
}

// add separated text and recording in wordList(word, isFiltered), wordLocation(wordID, urlID, index), linkWord(wordID, urlID)
func (s *CrawlerService) indexingAndRecordingWords(context context.Context, words []string, urlID int) error {
	countWords := 0
	countFilteredWords := 0
	re := regexp.MustCompile(`^\d+$`)

	for i, word := range words {
		needFiltered := re.MatchString(word)
		if needFiltered {
			countFilteredWords++
		}
		countWords++

		wordID, err := s.repository.AddWordList(context, &entity.WordList{Word: word, IsFiltred: needFiltered})
		if err != nil {
			return err
		}

		_, err = s.repository.AddWordLocation(context, &entity.WordLocation{WordID: wordID, URLID: urlID, Location: i})
		if err != nil {
			return err
		}

		_, err = s.repository.AddLinkWord(context, &entity.LinkWord{WordID: wordID, LinkID: urlID})
		if err != nil {
			return err
		}
	}

	return nil
}