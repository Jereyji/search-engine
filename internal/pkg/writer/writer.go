package writer

import (
	"fmt"
)

type Writer struct {
	ch chan []byte
}

func NewWriter() *Writer {
	return &Writer{
		ch: make(chan []byte),
	}
}

func (w *Writer) Write(value []byte) {
	w.ch <- value
}

func (w *Writer) Run() {
	for {
		resp, ok := <-w.ch
		if !ok {
			return
		}
		fmt.Println(string(resp))
	}
}

func (w *Writer) Close() {
	close(w.ch)
}