package montecarlo

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func ExampleSimulation() {
	tossEqualToOne := func() bool {
		// Simulate the toss of a six-sided die
		roll := rand.Intn(6) + 1

		if roll != 1 {
			return false
		}
		return true
	}

	// Used to get diferent random numbers on each execution
	rand.Seed(time.Now().UnixNano())

	const n = 100_000
	estimatedProbability := Simulation(tossEqualToOne, n)
	fmt.Printf("%.2f\n", estimatedProbability)
	// Output: 0.17
}

func tossEqualToSix() bool {
	// Simulate the toss of a six-sided die
	roll := rand.Intn(6) + 1

	if roll != 6 {
		return false
	}
	return true
}

func TestDieSimulation(t *testing.T) {
	const numSimsTest = 5_000_000
	const tolerance = 0.001

	expectedProbability := float64(1) / float64(6)
	actualProbability := Simulation(tossEqualToSix, numSimsTest)

	if !compareWithTolerance(expectedProbability, actualProbability, tolerance) {
		t.Errorf("expectedProbability=%f and actualProbability=%f => Not same by tolerance=%f\n", expectedProbability, actualProbability, tolerance)
	}

}

// Returns true if the numbers are equal with a tolerace, false otherwise
func compareWithTolerance(a, b, tolerance float64) bool {
	if diff := math.Abs(a - b); diff < tolerance {
		return true
	}

	return false
}

const (
	numsSimBenchmark       = 1_000_000
	numGoroutinesBenckmark = 10
)

func BenchmarkSimulationSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simulationSequential(tossEqualToSix, numsSimBenchmark)
	}
}

func BenchmarkSimulationConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simulationConcurrent(tossEqualToSix, numsSimBenchmark, numGoroutinesBenckmark)
	}
}
