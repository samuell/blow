// Flow process component that reads a line, line by line, into
// byte arrays.
// Author: Samuel Lampa
// Created: 2013-08-06

package blow

import (
	"bufio"
	"github.com/trustmaster/goflow"
	"log"
	"os"
)

type FileReader struct {
	flow.Component
	FileName <-chan string // Input port
	Line     chan<- []byte // Output port
}

func (fr *FileReader) OnFileName(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	} else {
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			// Create a copy of the current buffer slice to avoid data races
			lineCopy := append([]byte(nil), scan.Bytes()...)
			// Send the copy on the output port / channel
			fr.Line <- lineCopy
		}
		file.Close()
	}
}
