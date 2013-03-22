package genetic

import (
	"math"
	"math/rand"
	"bytes"
	"fmt"
	"strings"
	"time"
	"github.com/Tanner/Brainfuck-Go/src/brainfuck"
)

const numberGeneValues int = 8

type Entity struct {
	Genome []int
}

// NewEntity returns a new Entity with a random genome
func NewEntity(numberGenes int) *Entity {
	genome := make([]int, numberGenes, numberGenes)

	rand.Seed(time.Now().UTC().UnixNano())

	for i, _ := range genome {
		genome[i] = rand.Intn(numberGeneValues - 1)
	}

	e := Entity{genome}

	return &e
}

// NewEntityFromCode returns a new Entity from the given brainfuck code
func NewEntityFromCode(code string) *Entity {
	genome := make([]int, len(code), len(code))

	for i, v := range code {
		geneValue := 0

		switch(v) {
		case '>':
			geneValue = 0
		case '<':
			geneValue = 1
		case '+':
			geneValue = 2
		case '-':
			geneValue = 3
		case '.':
			geneValue = 4
		case ',':
			geneValue = 5
		case '[':
			geneValue = 6
		case ']':
			geneValue = 7
		}

		genome[i] = geneValue
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
func Crossover(e1 *Entity, e2 *Entity) (*Entity, *Entity, error) {
	if len(e1.Genome) != len(e2.Genome) {
		return nil, nil, fmt.Errorf("Parent genome lengths not the same, given %d, %d", len(e1.Genome), len(e2.Genome))
	}

	numberGenes := len(e1.Genome)

	halfwayIndex := (int)(numberGenes / 2)
	odd := 0

	if numberGenes % 2 != 0 {
		odd = 1
	}

	var e3, e4 Entity
	e3.Genome = make([]int, numberGenes, numberGenes)
	e4.Genome = make([]int, numberGenes, numberGenes)

	for i := 0; i < halfwayIndex + odd; i++ {
		e3.Genome[i] = e1.Genome[i]
		e3.Genome[halfwayIndex + i] = e2.Genome[halfwayIndex + i]

		e4.Genome[i] = e2.Genome[i]
		e4.Genome[halfwayIndex + i] = e1.Genome[halfwayIndex + i]
	}

	return &e3, &e4, nil
}

// Fitness rates the output of an entity's code output to the desired output
func (e *Entity) Fitness(in, correctOutput string, maxCycles int) float64 {
	output := new(bytes.Buffer)
	input := strings.NewReader(in)

	err := brainfuck.Run(e.Code(), output, input, maxCycles)

	if err != nil {
		// Code did not validate, give a low fitness
		return 0.0
	}

	outputString := output.String()
	fitness := 0.0

	correctOutputLength := len(correctOutput)

	for i := 0; i < len(outputString) && i < len(correctOutput); i++ {
		difference := math.Abs((float64) (correctOutput[i] - outputString[i]))

		fitness += (1.0 - difference / 255.0) * (1.0 / float64(correctOutputLength))
	}

	return fitness
}