package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Jereyji/search-engine/internal/pkg/request"
	"github.com/Jereyji/search-engine/internal/pkg/writer"
)

func (h *CrawlerHandler) PopularWords(ctx context.Context, w *writer.Writer, req *request.Request) {
	res, err := h.crawlerService.PopularWords(ctx)
	if err != nil {
		w.Write(InternalError(err))
		log.Println(err)
		return
	}

	output, err := json.Marshal(res)
	if err != nil {
		w.Write(InternalError(err))
		log.Println(err)
		return
	}

	w.Write(output)
}
