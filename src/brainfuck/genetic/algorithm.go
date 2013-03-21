package genetic

type Member struct {
	entity Entity
	fitness float64
}

type Algorithm struct {
	Population []Member
	GoalOutput string
	Input string
}

func NewAlgorithm(populationSize, numberGenes int, goalOutput, input string) *Algorithm {
	algorithm := new(Algorithm)

	algorithm.GoalOutput = goalOutput
	algorithm.Input = input
	algorithm.Population = make([]Member, populationSize, populationSize)

	for i, _ := range algorithm.Population {
		algorithm.Population[i].entity = *NewEntity(numberGenes)
	}

	return algorithm
}

// updateFitness updates the fitness stored in each Member of the Algorithm's Population
func (algorithm *Algorithm) updateFitness() {
	for _, member := range algorithm.Population {
		member.fitness = member.entity.Fitness(algorithm.Input, algorithm.GoalOutput)
	}
}