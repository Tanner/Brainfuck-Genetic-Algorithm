package genetic

import "testing"

func TestMutation(t *testing.T) {
	e := NewEntity(10)

	initialCode := e.Code()

	if	err := e.Mutate(0.9); err != nil {
		t.Error(err)
	}

	mutatedCode := e.Code()

	if initialCode == mutatedCode {
		t.Error("Entity's code did not mutate")
	}

	if err := e.Mutate(-1.3); err == nil {
		t.Error("Negative mutation rate should not be accepted")
	}

	if err := e.Mutate(1.3); err == nil {
		t.Error("Mutation rate greater than 1.0 should not be accepted")
	}
}

func TestCrossover(t *testing.T) {
	helper := func(numberGenes int) {
		e1 := NewEntity(10)
		e2 := NewEntity(10)

		e3, e4, err := Crossover(e1, e2)

		if err != nil {
			t.Error(err)
		}

		e1Code := e1.Code()
		e2Code := e2.Code()
		e3Code := e3.Code()
		e4Code := e4.Code()

		halfwayIndex := (int)(len(e1.Code()) / 2)

		if e1Code[:halfwayIndex] != e3Code[:halfwayIndex] || e2Code[halfwayIndex:] != e3Code[halfwayIndex:] {
			t.Error("Child from crossover did not get half genes from both parents")
		}

		if e2Code[:halfwayIndex] != e4Code[:halfwayIndex] || e1Code[halfwayIndex:] != e4Code[halfwayIndex:] {
			t.Error("Child from crossover did not get half genes from both parents")
		}
	}

	// Test odd number of genes
	helper(11)
	helper(3)

	// Test even number of genes
	helper(10)
	helper(4)
}

func TestFitness(t *testing.T) {
	e := NewEntityFromCode("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.")

	fitness := e.Fitness("", "Hello World!")

	if fitness != 1.0 {
		t.Errorf("Correct code did not get perfect fitness with correct output - got %f, want %f", fitness, 1.0)
	}
}