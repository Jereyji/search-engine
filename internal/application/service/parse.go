package service

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Document struct {
	doc *goquery.Document
}

func newDoc(link string) (*Document, error) {
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

	return &Document{doc}, nil
}

func (d Document) parseTitle(titleTextTag, titleLinkTag string) []TitleInfo {
	var res []TitleInfo

	d.doc.Find(titleLinkTag).Each(func(i int, selection *goquery.Selection) {
		link, exists := selection.Attr("href")
		if !exists {
			return
		}

		linkText := selection.Find(titleTextTag).Text()

		res = append(res, TitleInfo{link, linkText})
	})

	return res
}

func (d Document) parseArticle(articleTextTag, articleLinkTag string) []ArticleInfo {
	var res []ArticleInfo

	d.doc.Find(articleTextTag).Each(func(i int, selection *goquery.Selection) {
		mainText := selection.Text()
		res = append(res, ArticleInfo{MainText: mainText})

		selection.Find(articleLinkTag).Each(func(j int, linkSelection *goquery.Selection) {
			link, exists := linkSelection.Attr("href")
			if !exists {
				return
			}

			linkText := linkSelection.Text()
			
			res[i].RelatedLinks = append(res[i].RelatedLinks, link)
			res[i].RelatedLinkTexts = append(res[i].RelatedLinkTexts, linkText)
		})
	})

	return res
}
