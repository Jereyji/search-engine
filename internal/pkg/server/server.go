package server

import (
	"context"
	"log"

	router "github.com/Jereyji/search-engine/internal/pkg/console_router"
	"github.com/Jereyji/search-engine/internal/pkg/reader"
	"github.com/Jereyji/search-engine/internal/pkg/request"
	"github.com/Jereyji/search-engine/internal/pkg/writer"
)

func ListenAndServe(ctx context.Context, router *router.Router) {
	readCh := reader.NewReader()
	go readCh.Run(ctx)

	writeCh := writer.NewWriter()
	go writeCh.Run()
	defer writeCh.Close()

	for {
		select {
		case input, ok := <-readCh.Read():
			if !ok {
				return
			}

			req, err := request.ParseRequest(input)
			if err != nil {
				log.Printf("error during parsing request: %s", err)
				continue
			}

			if err = router.ServeConsole(writeCh, req); err != nil {
				log.Println(err)
			}
		case <-ctx.Done():
			return
		}
	}
}
