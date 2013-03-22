package main

import (
	"fmt"
	"bytes"
	"strings"
	"brainfuck"
	"brainfuck/genetic"
)

func main() {
	algorithm := genetic.NewAlgorithm(10, 20, 0.25, "Hi", "", 10000)

	for i := 0; i < 1000; i++ {
		fmt.Printf("Generation %d\n", algorithm.Generations)

		algorithm.Evolve()

		fmt.Printf("Best performing member (fitness %f)\n", algorithm.BestMember.Fitness)
		fmt.Printf("Code:   %s\n", algorithm.BestMember.Entity.Code())

		output := new(bytes.Buffer)
		input := strings.NewReader("")

		brainfuck.Run(algorithm.BestMember.Entity.Code(), output, input, 10000)
		fmt.Printf("Output: %s\n\n", output.String())
	}
}