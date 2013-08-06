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
			fr.Line <- scan.Bytes()
		}
	}
}
