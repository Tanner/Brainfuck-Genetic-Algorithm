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
	Generations int
	MutationRate float32
}

func NewAlgorithm(populationSize, numberGenes int, mutationRate float32, goalOutput, input string) *Algorithm {
	algorithm := new(Algorithm)

	algorithm.GoalOutput = goalOutput
	algorithm.Input = input
	algorithm.Population = make([]Member, populationSize, populationSize)
	algorithm.Generations = 0
	algorithm.MutationRate = mutationRate

	for i, _ := range algorithm.Population {
		algorithm.Population[i].entity = *NewEntity(numberGenes)
	}

	return algorithm
}

// Evolve evolves the entire population (crossover and mutates)
func (algorithm *Algorithm) Evolve() {
	NextPopulation := make([]Member, 0, len(algorithm.Population))

	for i := 0; i < cap(NextPopulation); i++ {
		parentA := algorithm.Select()
		parentB := algorithm.Select()

		childA, _, _ := Crossover(&parentA.entity, &parentB.entity)

		childA.Mutate(algorithm.MutationRate)

		newMember := Member{*childA, 0.0}

		NextPopulation = append(NextPopulation, newMember)
	}

	algorithm.Population = NextPopulation

	algorithm.Generations++
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
		if member.fitness == 0 || fitnessSum == 0 {
			normalizedFitness[i] = 0
		} else {
			normalizedFitness[i] = member.fitness / fitnessSum
		}
	}

	decision := rand.Float64()
	sum := 0.0

	for i, fitness := range normalizedFitness {
		sum += fitness

		if sum >= decision {
			return &algorithm.Population[i]
		}
	}

	return &algorithm.Population[rand.Intn(len(algorithm.Population))]
}

// updateFitness updates the fitness stored in each Member of the Algorithm's Population
func (algorithm *Algorithm) updateFitness() {
	for i, _ := range algorithm.Population {
		member := &algorithm.Population[i];

		member.fitness = member.entity.Fitness(algorithm.Input, algorithm.GoalOutput)
	}
}