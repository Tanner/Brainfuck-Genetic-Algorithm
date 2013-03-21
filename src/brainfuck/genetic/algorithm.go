package genetic

type Member struct {
	entity Entity
	fitness float32
}

type Algorithm struct {
	Population []Member
}

func NewAlgorithm(populationSize, numberGenes int) *Algorithm {
	algorithm := new(Algorithm)

	algorithm.Population = make([]Member, populationSize, populationSize)

	for i, _ := range algorithm.Population {
		algorithm.Population[i].entity = *NewEntity(numberGenes)
	}

	return algorithm
}