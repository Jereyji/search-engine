// package writer

// import (
// 	"fmt"
// )

// type Writer struct {
// 	ch chan []byte
// }

// func NewWriter() *Writer {
// 	return &Writer{
// 		ch: make(chan []byte),
// 	}
// }

// func (w *Writer) Write(value []byte) {
// 	w.ch <- value
// }

// func (w *Writer) Run() {
// 	for {
// 		resp, ok := <-w.ch
// 		if !ok {
// 			return
// 		}
// 		fmt.Println(string(resp))
// 	}
// }

// func (w *Writer) Close() {
// 	close(w.ch)
// }
package writer

import (
	"fmt"
	"os"
)

type Writer struct {
	ch     chan []byte
	file   *os.File
}

func NewWriter() *Writer {
	file, err := os.Create("result.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return nil
	}

	return &Writer{
		ch:   make(chan []byte),
		file: file,
	}
}

func (w *Writer) Write(value []byte) {
	w.ch <- value
}

func (w *Writer) Run() {
	defer w.file.Close() // Закрываем файл по окончании работы

	for {
		resp, ok := <-w.ch
		if !ok {
			return
		}

		// Запись в консоль
		fmt.Println(string(resp))

		// Добавление в result.json
		_, err := w.file.Write(resp)
		if err != nil {
			fmt.Println("haha")
		}
	}
}

func (w *Writer) Close() {
	close(w.ch)
}
