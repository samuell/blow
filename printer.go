// Flow process component that prints a byte array as a line of text
// to the terminal (stdout).
// Author: Samuel Lampa
// Created: 2013-08-06

package blow

import (
	"fmt"
	"github.com/trustmaster/goflow"
	"log"
	"sync"
)

type Printer struct {
	flow.Component
	Line      <-chan []byte // Input
	StateLock *sync.Mutex   // Thread-safety lock
	cnt       int
}

func NewPrinter() *Printer {
	p := new(Printer)
	p.StateLock = new(sync.Mutex)
	return p
}

// Prints a line when it gets it
func (p *Printer) OnLine(line []byte) {
	p.cnt++
	log.Printf("[pr][%d]: %s", p.cnt, line)
	fmt.Println(string(line))
}
