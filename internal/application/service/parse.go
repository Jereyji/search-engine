package service

import (
	"net/http"
	"net/url"
	"strings"

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

func (d Document) parseTitle(baseURL, titleTextTag, titleLinkTag string) []TitleInfo {
	var res []TitleInfo

	d.doc.Find(titleLinkTag).Each(func(i int, selection *goquery.Selection) {
		link, exists := selection.Attr("href")
		if !exists {
			return
		}

		if !isLinkValid(link) {
			link = setBaseURL(link, baseURL)
		}

		linkText := selection.Find(titleTextTag).Text()

		res = append(res, TitleInfo{link, linkText})
	})

	return res
}

func (d Document) parseArticle(baseURL, articleTextTag, articleLinkTag string) []ArticleInfo {
	var res []ArticleInfo

	d.doc.Find(articleTextTag).Each(func(i int, selection *goquery.Selection) {
		mainText := selection.Text()
		res = append(res, ArticleInfo{MainText: mainText})

		selection.Find(articleLinkTag).Each(func(j int, linkSelection *goquery.Selection) {
			link, exists := linkSelection.Attr("href")
			if !exists {
				return
			}

			if !isLinkValid(link) {
				link = setBaseURL(link, baseURL)
			}
			
			res[i].RelatedLinks = append(res[i].RelatedLinks, TitleInfo{link, linkSelection.Text()})
		})
	})

	return res
}

func isLinkValid(link string) bool {
	return strings.HasPrefix(link, "http")
}

func getBaseURL(link string) string {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return ""
	}

	baseURL := parsedURL.Scheme + "://" + parsedURL.Host
	return baseURL
}

func setBaseURL(link, baseURL string) string {
	if !isLinkValid(link) {
		return baseURL + link
	}
	return link
}