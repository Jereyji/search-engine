package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/Jereyji/search-engine/internal/application/service"
	"github.com/Jereyji/search-engine/internal/pkg/request"
	"github.com/Jereyji/search-engine/internal/pkg/writer"
)

type contextKey string

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
	log.Println(ctx)
	
	// depth, ok := req.GetValue(depthFlags).(int)
	// if !ok {
	// 	w.Write(BadRequest(depthFlags))
	// 	return
	// }

	key := contextKey("dataLinks")
	// dataLinks := ctx.Value(key).([]service.DataLinks)
	dataLinks := ctx.Value(key)
	fmt.Println(dataLinks)
	// dataLinks, ok := ctx.Value(key).([]service.DataLinks)
	// if !ok {
	// 	log.Println("don't found links : ", dataLinks)
	// 	return
	// }

	// if err := h.crawlerService.Crawl(ctx, dataLinks, 1); err != nil {
	// 	log.Println(err)
	// 	return
	// }

	log.Println("GOOD")
	// w.Write()
}