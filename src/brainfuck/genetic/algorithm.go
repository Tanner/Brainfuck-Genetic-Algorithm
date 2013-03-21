package genetic

import "math/rand"

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

// Select selects a member from the population using roulette wheel selection
func (algorithm *Algorithm) Select() *Member {
	normalizedFitness := make([]float64, len(algorithm.Population))
	fitnessSum := 0.0

	algorithm.updateFitness()

	for _, member := range algorithm.Population {
		fitnessSum += member.fitness
	}

	for i, member := range algorithm.Population {
		normalizedFitness[i] = member.fitness / fitnessSum
	}

	decision := rand.Float64()
	sum := 0.0

	for i, fitness := range normalizedFitness {
		sum += fitness

		if sum >= decision {
			return &algorithm.Population[i]
		}
	}

	return nil
}

// updateFitness updates the fitness stored in each Member of the Algorithm's Population
func (algorithm *Algorithm) updateFitness() {
	for _, member := range algorithm.Population {
		member.fitness = member.entity.Fitness(algorithm.Input, algorithm.GoalOutput)
	}
}