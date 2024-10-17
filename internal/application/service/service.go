package service

import (
	"fmt"
	"log"
	"net/http"

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

type Service struct {
	repository repository_interface.SearchRepository
}

func NewService(repository repository_interface.SearchRepository) *Service {
	return &Service{
		repository: repository,
	}
}

type DataLinks []struct {
	Url      string `yaml:"url"`
	Selector string `yaml:"selector"`
	Text     string `yaml:"text"`
}

func Crawl(dataLinks DataLinks, maxDepth int) {
	for _, data := range dataLinks {
		res, err := http.Get(data.Url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(data.Url) //

		doc.Find(data.Selector).Each(func(i int, s *goquery.Selection) {
			title := s.Find(data.Text).Text()

			href, exists := s.Attr("href")
			if exists {
				fmt.Printf("Article %d:\nTitle: %s\nLink: %s\n", i, title, href)
			} else {
				fmt.Printf("Article %d: %s (no link found)\n", i, title)
			}
		})
	}
}
