// Flow process component that converts DNA or RNA sequences in FASTA
// format to their base complement.
// Author: Samuel Lampa
// Created: 2013-08-05

package blow

import (
	"github.com/trustmaster/goflow"
)

var baseConv = [256]byte{
	'A': 'T',
	'T': 'A',
	'C': 'G',
	'G': 'C',
	'N': 'N',
}

type BaseComplementer struct {
	flow.Component                         // Embedding "superclass"
	Sequence                 <-chan []byte // Input port
	BaseComplementedSequence chan<- []byte // Output port
}

func (bc *BaseComplementer) OnSequence(sequence []byte) {
    // Copy the array
    sequence = append([]byte(nil), sequence...)
	for pos := range sequence {
		sequence[pos] = baseConv[sequence[pos]]
	}
	bc.BaseComplementedSequence <- sequence
}
