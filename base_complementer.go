// Flow process component that converts DNA or RNA sequences in FASTA
// format to their base complement.
// Author: Samuel Lampa
// Created: 2013-08-05

package blow

import (
	"github.com/trustmaster/goflow"
	"log"
)

var baseConv = [256]byte{
	'A': 'T',
	'T': 'A',
	'C': 'G',
	'G': 'C',
	'N': 'N',
	'0': '0',
	'1': '1',
	'2': '2',
	'3': '3',
	'4': '4',
	'5': '5',
	'6': '6',
	'7': '7',
	'8': '8',
	'9': '9',
	' ': ' ',
}

type BaseComplementer struct {
	flow.Component                         // Embedding "superclass"
	Sequence                 <-chan []byte // Input port
	BaseComplementedSequence chan<- []byte // Output port
	Name                     string
	cnt                      int
}

func NewBaseComplementer(name string) *BaseComplementer {
	bc := new(BaseComplementer)
	bc.Name = name
	return bc
}

func (bc *BaseComplementer) OnSequence(sequence []byte) {
	bc.cnt++
	log.Printf("[%s][%d]: %s", bc.Name, bc.cnt, sequence)
	// Copy the array
	sequence = append([]byte(nil), sequence...)
	for pos := range sequence {
		sequence[pos] = baseConv[sequence[pos]]
	}
	bc.BaseComplementedSequence <- sequence
}
