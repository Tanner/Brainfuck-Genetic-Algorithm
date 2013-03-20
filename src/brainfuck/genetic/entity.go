package genetic

import (
	"math/rand"
)

const numberGenes int = 100
const numberGeneValues int = 8

type Entity struct {
	Genome [numberGenes]int
}

func NewEntity() *Entity {
	var genome [numberGenes]int

	for i, _ := range genome {
		genome[i] = rand.Intn(numberGeneValues + 1)
	}

	e := Entity{genome}

	return &e
}