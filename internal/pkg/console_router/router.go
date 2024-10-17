package router

import (
	"fmt"

	"github.com/Jereyji/search-engine/internal/pkg/request"
	"github.com/Jereyji/search-engine/internal/pkg/writer"
)

type HandlerFunc = func(w *writer.Writer, req *request.Request)

type Router struct {
	routes map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}

func (r *Router) HandleFunc(command string, hFunc HandlerFunc) {
	r.routes[command] = hFunc
}

func (r Router) ServeConsole(writeCh *writer.Writer, req *request.Request) error {
	if cmd, exists := r.routes[req.GetCommand()]; exists {
		cmd(writeCh, req)
	} else {
		return fmt.Errorf("unknown command: %s", req.GetCommand())
	}

	return nil
}
