package genetic

import (
	"math/rand"
	"bytes"
)

const numberGenes int = 100
const numberGeneValues int = 8

const mutationRate float32 = 0.25

type Entity struct {
	Genome [numberGenes]int
}

// NewEntity returns a new Entity with a random genome
func NewEntity() *Entity {
	var genome [numberGenes]int

	for i, _ := range genome {
		genome[i] = rand.Intn(numberGeneValues - 1)
	}

	e := Entity{genome}

	return &e
}

// Code returns the brainfuck code for the given Entity's genome
func (e *Entity) Code() string {
	var code bytes.Buffer
	var instruction rune

	for _, v := range e.Genome {
		switch(v) {
		case 0:
			instruction = '>'
		case 1:
			instruction = '<'
		case 2:
			instruction = '+'
		case 3:
			instruction = '-'
		case 4:
			instruction = '.'
		case 5:
			instruction = ','
		case 6:
			instruction = '['
		case 7:
			instruction = ']'
		}

		code.WriteRune(instruction)
	}

	return code.String()
}

// Mutate randomly mutates random genes according to the mutate rate constant
func (e *Entity) Mutate() {
	for i, _ := range e.Genome {
		if rand.Float32() <= mutationRate {
			e.Genome[i] = rand.Intn(numberGeneValues - 1)
		}
	}
}