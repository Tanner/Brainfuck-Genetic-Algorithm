package genetic

import (
	"math/rand"
	"bytes"
	"fmt"
)

const numberGenes int = 100
const numberGeneValues int = 8

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

// Mutate randomly mutates random genes according to the mutate rate
func (e *Entity) Mutate(mutationRate float32) error {
	if mutationRate < 0 || mutationRate > 1 {
		return fmt.Errorf("genetic: mutate rate not between [0.0, 1.0] received %f", mutationRate)
	}

	for i, _ := range e.Genome {
		if rand.Float32() <= mutationRate {
			e.Genome[i] = rand.Intn(numberGeneValues - 1)
		}
	}

	return nil
}

// Crossover two entities' genomes at their halfway point, returning two new children entities
func Crossover(e1 *Entity, e2 *Entity) (*Entity, *Entity) {
	halfwayIndex := (int)(numberGenes / 2)
	odd := 0

	if numberGenes % 2 != 0 {
		odd = 1
	}

	var e3, e4 Entity

	for i := 0; i < halfwayIndex + odd; i++ {
		e3.Genome[i] = e1.Genome[i]
		e3.Genome[halfwayIndex + i] = e2.Genome[halfwayIndex + i]

		e4.Genome[i] = e2.Genome[i]
		e4.Genome[halfwayIndex + i] = e1.Genome[halfwayIndex + i]
	}

	return &e3, &e4
}