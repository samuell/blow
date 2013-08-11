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
	cnt      int
}

func (fr *FileReader) OnFileName(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	} else {
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			line := scan.Bytes()
			fr.cnt++
			log.Printf("[fr][%d]: %s", fr.cnt, line)
			fr.Line <- line
		}
	}
}
