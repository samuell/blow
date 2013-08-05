// Flow process component that converts DNA or RNA sequences in FASTA
// format to their base complement.
// Author: Samuel Lampa
// Created: 2013-08-05

package blow

import (
	"github.com/trustmaster/goflow"
)

type BaseComplementer struct {
	flow.Component                         // Embedding "superclass"
	Sequence                 <-chan []byte // Input port
	BaseComplementedSequence chan<- []byte // Output port
}

func (bc *BaseComplementer) OnSequence(sequence []byte) {
	for pos := range sequence {
		if sequence[pos] == 'A' {
			sequence[pos] = 'T'
		} else if sequence[pos] == 'T' {
			sequence[pos] = 'A'
		} else if sequence[pos] == 'C' {
			sequence[pos] = 'G'
		} else if sequence[pos] == 'G' {
			sequence[pos] = 'C'
		}
	}
	bc.BaseComplementedSequence <- sequence
}
