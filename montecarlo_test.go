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
		t.Errorf("expectedProbability=%f and actualProbability=%f => Not same Even by Tolerance\n", expectedProbability, actualProbability)
	}

}

// Returns true if the numbers are equal with a tolerace, false otherwise
func compareWithTolerance(a, b float64) bool {
	tolerance := 0.001
	if diff := math.Abs(a - b); diff < tolerance {
		return true
	}

	return false
}
