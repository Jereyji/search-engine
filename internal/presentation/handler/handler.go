package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/Jereyji/search-engine/internal/application/service"
	"github.com/Jereyji/search-engine/internal/pkg/request"
	"github.com/Jereyji/search-engine/internal/pkg/writer"
)

type CrawlerHandler struct {
	crawlerService *service.CrawlerService
}

func NewCrawlerHandler(crawlerService *service.CrawlerService) *CrawlerHandler {
	return &CrawlerHandler{
		crawlerService: crawlerService,
	}
}

const (
	depthFlags = "--depth"
)

func (h *CrawlerHandler) Crawl(ctx context.Context, w *writer.Writer, req *request.Request) {	
	depth, ok := req.GetValue(depthFlags).(string)
	if !ok {
		w.Write(BadRequest(depthFlags))
		return
	}
	fmt.Println(depth)

	// dataLink, ok := req.GetValue(depthFlags).(string)
	// if !ok {
	// 	w.Write(BadRequest(depthFlags))
	// 	log.Println("don't found links : ", dataLink)
	// 	return
	// }

	dataLink := service.DataLinks {
		Url: "https://www.gazeta.ru/",
		Selector: "a.b_ear",
		Text: "div.b_ear-title",
	}
	
	if err := h.crawlerService.Crawl(ctx, dataLink, 1); err != nil {
		log.Println(err)
		return
	}

	log.Println("GOOD")
	// w.Write()
}