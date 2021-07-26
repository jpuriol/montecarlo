package montecarlo

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func tossSixSidedDie() int {
	return rand.Intn(6) + 1
}

func tossEqualToSix() bool {
	if tossSixSidedDie() == 6 {
		return true
	}

	return false
}

func TestDieSimulation(t *testing.T) {
	// Get diferent random numbers on each exection
	rand.Seed(time.Now().UnixNano())

	expectedProbability := float64(1) / float64(6)
	actualProbability := Simulation(tossEqualToSix)

	if !compareWithTolerance(expectedProbability, actualProbability) {
		t.Errorf("expectedProbability=%f and actualProbability=%f => Not same by tolerance=%f\n", expectedProbability, actualProbability, tolerance)
	}

}

func BenchmarkSimulationSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simulationSequential(tossEqualToSix)
	}
}

func BenchmarkSimulationConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simulationConcurrent(tossEqualToSix)
	}
}

const tolerance = 0.0001

// Returns true if the numbers are equal with a tolerace, false otherwise
func compareWithTolerance(a, b float64) bool {
	if diff := math.Abs(a - b); diff < tolerance {
		return true
	}

	return false
}
