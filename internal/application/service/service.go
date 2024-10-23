package service

import (
	"context"
	// "fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Jereyji/search-engine/internal/domain/entity"
	"github.com/Jereyji/search-engine/internal/domain/repository_interface"
	"github.com/PuerkitoBio/goquery"
)

const (
	linkWord    = true
	notLinkWord = false
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

type DataURL struct {
	URL             string `yaml:"url"`
	TitleSelector   string `yaml:"title_selector"`
	TitleText       string `yaml:"title_text"`
	ArticleSelector string `yaml:"article_selector"`
	ArticleLink     string `yaml:"article_link"`
}

type Response struct {
	URL                string
	CountWords         int
	CountFilteredWords int
}

// Загрузка страницы из списка обхода;
// Проверка, есть ли такая страница уже в индексе, если страницы в индексе нет, то она добавляется;
// Выделение новые ссылки со страницы и добавляются в список обхода;
// Далее паука переходит к следующему документу.

func (s *CrawlerService) Crawl(ctx context.Context, data DataURL, depth int) (*Response, error) {
	res := Response{
		URL: data.URL,
	}

	for i := 0; i < depth; i++ {
		tempURL, err := s.repository.GetURL(ctx, data.URL)
		if err != nil {
			return nil, err
		}

		if tempURL != nil {
			continue
		}

		URL := entity.URLList{
			Link: data.URL,
		}

		URL.ID, err = s.repository.AddURL(ctx, &URL)
		if err != nil {
			return nil, err
		}

		doc, err := newDoc(URL.Link)
		if err != nil {
			return nil, err
		}
		// func work with Articles
		if res.CountWords, res.CountFilteredWords, err = s.parseArticle(ctx, doc, &URL, data.ArticleSelector); err != nil {
			return nil, err
		}

		// func work with Titles
	}

	return &res, nil
}

func (s *CrawlerService) parseArticle(ctx context.Context, doc *goquery.Document, URL *entity.URLList, selector string) (int, int, error) {
	var (
		countFilteredWords int
		countWords         int
		err                error
	)

	doc.Find(selector).Each(func(i int, selection *goquery.Selection) {
		paragraphText := selection.Text()

		countWords, countFilteredWords, err = s.addText(ctx, paragraphText, URL.ID, notLinkWord)
		if err != nil {
			return
		}

		selection.Find("a").Each(func(j int, linkSelection *goquery.Selection) {
			link, exists := linkSelection.Attr("href")
			if !exists {
				return
			}

			relatedURL, err := s.addRelatedURL(ctx, URL.ID, link)
			if err != nil {
				return
			}

			linkText := linkSelection.Text()
			tempCountWords, tempFilteredWords, err := s.addText(ctx, linkText, relatedURL.ID, linkWord)
			if err != nil {
				return
			}

			countFilteredWords += tempFilteredWords
			countWords += tempCountWords
		})
	})

	return countWords, countFilteredWords, err
}

func (s *CrawlerService) addRelatedURL(ctx context.Context, fromURLID int, link string) (*entity.URLList, error) {
	tempURL, err := s.repository.GetURL(ctx, link)
	if err != nil {
		return nil, err
	}

	if tempURL != nil {
		return tempURL, nil
	}

	URL := entity.URLList{
		Link: link,
	}

	URL.ID, err = s.repository.AddURL(ctx, &URL)
	if err != nil {
		return nil, err
	}

	_, err = s.repository.AddLinkBetweenURLs(ctx, &entity.LinkBetweenURL{FromURLID: fromURLID, ToURLID: URL.ID})
	if err != nil {
		return nil, err
	}

	return &URL, nil
}

// add separated text and recording in wordList(word, isFiltered), wordLocation(wordID, urlID, index), linkWord(wordID, urlID)
func (s *CrawlerService) addText(ctx context.Context, text string, urlID int, linked bool) (int, int, error) {
	re := regexp.MustCompile(`^\d+$`)
	countFilteredWords := 0
	countWords := 0

	words := strings.Fields(text)
	if words == nil {
		return 0, 0, nil
	}

	for i, word := range words {
		needFiltered := re.MatchString(word)
		if needFiltered {
			countFilteredWords++
		}
		countWords++

		wordID, err := s.repository.AddWordList(ctx, &entity.WordList{Word: word, IsFiltred: needFiltered})
		if err != nil {
			return 0, 0, err
		}

		_, err = s.repository.AddWordLocation(ctx, &entity.WordLocation{WordID: wordID, URLID: urlID, Location: i})
		if err != nil {
			return 0, 0, err
		}

		if linked {
			_, err = s.repository.AddLinkWord(ctx, &entity.LinkWord{WordID: wordID, LinkID: urlID})
			if err != nil {
				return 0, 0, err
			}
		}
	}

	return countWords, countFilteredWords, nil
}

func newDoc(link string) (*goquery.Document, error) {
	httpResponse, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// foundLink, exists := selection.Attr("href")
// if !exists {
// 	return
// }

// _, err = s.addRelatedURL(ctx, URL.ID, foundLink)
// if err != nil {
// 	return
// }

// func (s *CrawlerService) addURL(ctx context.Context, link string) (*entity.URLList, error) {
// 	_, err := s.repository.GetURL(ctx, link)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &curURL, nil
// }