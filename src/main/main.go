package main

import (
	"brainfuck"
	"brainfuck/genetic"
	"bytes"
	"fmt"
	"strings"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	algorithm := genetic.NewAlgorithm(500, 10, 0.5, "Hi", "", 1000)

	desiredFitness := 0.8
	bestFitness := 0.0

	for bestFitness < desiredFitness {
		now := time.Now()

		fmt.Printf("Generation %d", algorithm.Generations)

		algorithm.Evolve()

		fmt.Printf(" (%s)\n", time.Since(now))

		bestFitness = algorithm.BestMember.Fitness

		fmt.Printf("Best performing member (fitness %f)\n", bestFitness)
		fmt.Printf("Code:   %s\n", algorithm.BestMember.Entity.Code())

		output := new(bytes.Buffer)
		input := strings.NewReader("")

		brainfuck.Run(algorithm.BestMember.Entity.Code(), output, input, 1000)
		fmt.Printf("Output: %s\n\n", output.String())
	}
}
