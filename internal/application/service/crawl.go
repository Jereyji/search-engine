package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Jereyji/search-engine/internal/domain/entity"
)

var ErrExistRow = errors.New("row already exists")

const (
	parsed    = true
	notParsed = false
)

type DataURL struct {
	URL            string `yaml:"url"`
	TitleTextTag   string `yaml:"title_text_tag"`
	TitleLinkTag   string `yaml:"title_link_tag"`
	ArticleTextTag string `yaml:"article_text_tag"`
	ArticleLinkTag string `yaml:"article_link_tag"`
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
		for ; j < sizeLinksList; j++ {
			URL, err := s.ensureURLExists(ctx, linksList[j])
			if err != nil {
				return nil, err
			}

			if URL.Is_parsed {
				continue
			}

			doc, err := newDoc(linksList[j])
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
				fmt.Println("Title - ", *curRes) // Для наглядности
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
				fmt.Println("Article - ", curRes) // Для наглядности
			}

			URL.ChangeParseStatus(parsed)
			if err := s.repository.UpdateURL(ctx, URL); err != nil {
				return nil, err
			}
		}
	}

	return res, nil
}

func (s *CrawlerService) ensureURLExists(ctx context.Context, link string) (*entity.URLList, error) {
	tempURL, err := s.repository.URL(ctx, link)
	if err != nil {
		return nil, err
	}

	if tempURL.ID != 0 {
		return tempURL, nil
	}

	URL := entity.URLList{
		Link:      link,
		Is_parsed: notParsed,
	}

	URL.ID, err = s.repository.CreateURL(ctx, &URL)
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

		wordID, err := s.repository.CreateWord(ctx, &entity.WordList{Word: word, IsFiltred: needFiltered})
		if err != nil {
			return 0, 0, err
		}

		_, err = s.repository.CreateWordLocation(ctx, &entity.WordLocation{WordID: wordID, URLID: urlID, Location: i})
		if err != nil {
			return 0, 0, err
		}

		if linked != 0 {
			_, err = s.repository.CreateLinkWord(ctx, &entity.LinkWord{WordID: wordID, LinkID: linked})
			if err != nil {
				return 0, 0, err
			}
		}
	}

	return countWords, countFilteredWords, nil
}
