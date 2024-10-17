package reader

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
)

type Reader struct {
	ch chan string
}

func NewReader() *Reader {
	return &Reader{
		ch: make(chan string),
	}
}

func (r *Reader) Run(ctx context.Context) {
	reader := bufio.NewReader(os.Stdin)
	defer close(r.ch)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Error reading input: %v\n", err)
				return
			}
			input = strings.TrimSpace(input)
			r.ch <- input
		}
	}
}

func (r *Reader) Read() <-chan string {
	return r.ch
}
