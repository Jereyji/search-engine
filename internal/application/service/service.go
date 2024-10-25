package service

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/Jereyji/search-engine/internal/domain/entity"
	"github.com/Jereyji/search-engine/internal/domain/repository_interface"
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

const (
	errExistRow = "row is exist"
	linkWord    = true
	notLinkWord = false
)

type CrawlerService struct {
	repository repository_interface.CrawlerRepository
}

func NewCrawlerService(repository repository_interface.CrawlerRepository) *CrawlerService {
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

func (s *CrawlerService) Crawl(ctx context.Context, data DataURL, depth int) ([]Response, error) {
	res := []Response{
		{URL: data.URL},
	}

	dataURLList := []DataURL{data}

	for i := 0; i < depth; i++ {
		for j, data := range dataURLList {
			URL, err := s.ensureURLExists(ctx, data.URL)
			if err != nil {
				return nil, err
			}

			if errors.Is(err, errors.New(errExistRow)) {
				continue
			}

			doc, err := newDoc(URL.Link)
			if err != nil {
				return nil, err
			}

			parsedTitles := doc.parseTitle(data.TitleTextTag, data.TitleLinkTag)
			for _, title := range parsedTitles {
				curRes, err := title.storeToRepository(ctx, s, URL)
				if err != nil {
					return nil, err
				}

				// dataURLList = append(dataURLList, curRes.LinkList)
				res = append(res,  *curRes)
			}

			parsedArticles := doc.parseArticle(data.ArticleTextTag, data.ArticleLinkTag)
			for _, article := range parsedArticles {
				curRes, err := article.storeToRepository(ctx, s, URL)
				if err != nil {
					return nil, err
				}

				// dataURLList = append(dataURLList, curRes.LinkList)
				res = append(res,  *curRes)
			}

			dataURLList = dataURLList[j:]
		}
	}

	return res, nil
}

func (s *CrawlerService) ensureURLExists(ctx context.Context, link string) (*entity.URLList, error) {
	tempURL, err := s.repository.GetURL(ctx, link)
	if err != nil {
		return nil, err
	}

	if tempURL != nil {
		return tempURL, errors.New(errExistRow)
	}

	URL := entity.URLList{
		Link: link,
	}

	URL.ID, err = s.repository.AddURL(ctx, &URL)
	if err != nil {
		return nil, err
	}

	return &URL, nil
}

// add separated text and recording in wordList(word, isFiltered), wordLocation(wordID, urlID, index), linkWord(wordID, urlID)
func (s *CrawlerService) addText(ctx context.Context, text string, urlID int, linked int) (int, int, error) {
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

		if linked != 0 {
			_, err = s.repository.AddLinkWord(ctx, &entity.LinkWord{WordID: wordID, LinkID: linked})
			if err != nil {
				return 0, 0, err
			}
		}
	}

	return countWords, countFilteredWords, nil
}
