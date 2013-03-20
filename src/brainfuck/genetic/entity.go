package genetic

import (
	"math/rand"
)

const numberGenes int = 8

type Entity struct {
	genome [numberGenes]int
}

func NewEntity() *Entity {
	var genome [numberGenes]int

	for i, _ := range genome {
		genome[i] = rand.Intn(numberGenes + 1)
	}

	e := Entity{genome}

	return &e
}