package genetic

import (
	"math/rand"
	"time"
)

type Member struct {
	Entity Entity
	Fitness float64
}

type Algorithm struct {
	Population []Member
	GoalOutput string
	Input string
	Generations int
	MutationRate float32
	BestMember *Member
}

func NewAlgorithm(populationSize, numberGenes int, mutationRate float32, goalOutput, input string) *Algorithm {
	algorithm := new(Algorithm)

	algorithm.GoalOutput = goalOutput
	algorithm.Input = input
	algorithm.Population = make([]Member, populationSize, populationSize)
	algorithm.Generations = 0
	algorithm.MutationRate = mutationRate
	algorithm.BestMember = nil

	for i, _ := range algorithm.Population {
		algorithm.Population[i].Entity = *NewEntity(numberGenes)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	return algorithm
}

// Evolve evolves the entire population (crossover and mutates)
func (algorithm *Algorithm) Evolve() {
	NextPopulation := make([]Member, 0, len(algorithm.Population))

	for i := 0; i < cap(NextPopulation); i++ {
		parentA := algorithm.Select()
		parentB := algorithm.Select()

		childA, _, _ := Crossover(&parentA.Entity, &parentB.Entity)

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
		fitnessSum += member.Fitness
	}

	for i, member := range algorithm.Population {
		if member.Fitness == 0 || fitnessSum == 0 {
			normalizedFitness[i] = 0
		} else {
			normalizedFitness[i] = member.Fitness / fitnessSum
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

		member.Fitness = member.Entity.Fitness(algorithm.Input, algorithm.GoalOutput)

		if algorithm.BestMember == nil || member.Fitness > algorithm.BestMember.Fitness {
			algorithm.BestMember = member
		}
	}
}