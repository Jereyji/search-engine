package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Jereyji/search-engine/internal/domain/entity"
	"github.com/Jereyji/search-engine/internal/domain/repository"
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

var ErrExistRow = errors.New("row already exists")

const (
	parsed    = true
	notParsed = false
)

type CrawlerService struct {
	repository repository.CrawlerRepository
}

func NewCrawlerService(repository repository.CrawlerRepository) *CrawlerService {
	return &CrawlerService{
		repository: repository,
	}
}

type DataURL struct {
	URL            string `yaml:"url"`
	TitleTextTag   string `yaml:"title_text_tag"`
	TitleLinkTag   string `yaml:"title_link_tag"`
	ArticleTextTag string `yaml:"article_text_tag"`
	ArticleLinkTag string `yaml:"article_link_tag"`
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

func (s *CrawlerService) Crawl(ctx context.Context, depth int, data DataURL) ([]Response, error) {
	res := []Response{
		{URL: data.URL},
	}

	baseURL := getBaseURL(data.URL)
	linksList := []string{data.URL}

	for i, j := 0, 0; i < depth; i++ {
		sizeLinksList := len(linksList)
		// fmt.Println("cur size = ", j, sizeLinksList)
		for ; j < sizeLinksList; j++ {
			URL, err := s.ensureURLExists(ctx, linksList[j])
			if err != nil {
				return nil, err
			}

			if URL.Is_parsed {
				continue
			}

			doc, err := newDoc(URL.Link)
			if err != nil {
				return nil, err
			}

			parsedTitles := doc.parseTitle(baseURL, data.TitleTextTag, data.TitleLinkTag)
			for _, title := range parsedTitles {
				curRes, err := title.storeToRepository(ctx, s, URL)
				if err != nil {
					return nil, err
				}

				if curRes != nil {
					linksList = append(linksList, title.Link)
					res = append(res, *curRes)
				}
				fmt.Println("Title - ", *curRes)
			}

			parsedArticles := doc.parseArticle(baseURL, data.ArticleTextTag, data.ArticleLinkTag)
			for _, article := range parsedArticles {
				curRes, err := article.storeToRepository(ctx, s, URL)
				if err != nil {
					return nil, err
				}

				if curRes != nil {
					for _, relatedLink := range article.RelatedLinks {
						linksList = append(linksList, relatedLink.Link)
					}
					res = append(res, curRes...)
				}
				// fmt.Println("Article - ", curRes)
			}

			URL.ChangeParseStatus(parsed)
			if err := s.repository.URLList.Update(ctx, URL); err != nil {
				return nil, err
			}
		}
	}

	return res, nil
}

func (s *CrawlerService) ensureURLExists(ctx context.Context, link string) (*entity.URLList, error) {
	tempURL, err := s.repository.URLList.URL(ctx, link)
	if err != nil {
		return nil, err
	}

	if tempURL != nil {
		return tempURL, nil
	}

	URL := entity.URLList{
		Link:      link,
		Is_parsed: notParsed,
	}

	URL.ID, err = s.repository.URLList.Create(ctx, &URL)
	if err != nil {
		return nil, err
	}

	return &URL, nil
}

// add separated text and recording in wordList(word, isFiltered), wordLocation(wordID, urlID, index), linkWord(wordID, urlID)
func (s *CrawlerService) addText(ctx context.Context, text *string, urlID int, linked int) (int, int, error) {
	re := regexp.MustCompile(`^\d+$`)
	countFilteredWords := 0
	countWords := 0

	words := strings.Fields(*text)
	if words == nil {
		return 0, 0, nil
	}

	for i, word := range words {
		needFiltered := re.MatchString(word)
		if needFiltered {
			countFilteredWords++
		}
		countWords++

		wordID, err := s.repository.WordList.Create(ctx, &entity.WordList{Word: word, IsFiltred: needFiltered})
		if err != nil {
			return 0, 0, err
		}

		_, err = s.repository.WordLocation.Create(ctx, &entity.WordLocation{WordID: wordID, URLID: urlID, Location: i})
		if err != nil {
			return 0, 0, err
		}

		if linked != 0 {
			_, err = s.repository.LinkWord.Create(ctx, &entity.LinkWord{WordID: wordID, LinkID: linked})
			if err != nil {
				return 0, 0, err
			}
		}
	}

	return countWords, countFilteredWords, nil
}
