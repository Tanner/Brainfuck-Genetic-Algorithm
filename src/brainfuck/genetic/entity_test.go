package genetic

import "testing"

func TestMutation(t *testing.T) {
	e := NewEntity()

	initialCode := e.Code()

	if	err := e.Mutate(0.25); err != nil {
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