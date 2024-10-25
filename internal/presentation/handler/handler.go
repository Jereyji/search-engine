package handler

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/Jereyji/search-engine/internal/application/service"
	"github.com/Jereyji/search-engine/internal/pkg/config"
	"github.com/Jereyji/search-engine/internal/pkg/request"
	"github.com/Jereyji/search-engine/internal/pkg/writer"
)

type CrawlerHandler struct {
	crawlerService *service.CrawlerService
	config         *config.Config
}

func NewCrawlerHandler(crawlerService *service.CrawlerService, cfg *config.Config) *CrawlerHandler {
	return &CrawlerHandler{
		crawlerService: crawlerService,
		config:         cfg,
	}
}

const (
	depthFlags = "--depth"
)

func (h *CrawlerHandler) Crawl(ctx context.Context, w *writer.Writer, req *request.Request) {
	depthStr, ok := req.GetValue(depthFlags).(string)
	if !ok {
		w.Write(BadRequest(MissingFlag, depthFlags))
		return
	}

	depth, err := strconv.Atoi(depthStr)
	if err != nil {
		w.Write(BadRequest(IncorrectData, depthStr))
		return
	}

	var allResults []service.Response

	for _, dataLink := range h.config.DataURLs {
		res, err := h.crawlerService.Crawl(ctx, dataLink, depth)
		if err != nil {
			w.Write(InternalError(err))
			log.Println(err)
			return
		}
		allResults = append(allResults, res)
	}

	output, err := json.Marshal(allResults)
	if err != nil {
		w.Write(InternalError(err))
		log.Println(err)
		return
	}

	w.Write(output)
}
