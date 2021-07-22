// Package montecarlo provides a simple interface to execute Monte Carlo simulations.
package montecarlo

const numTrials = 1_000_000

// Returns the probability of the event of interest after running simulations
//
// The occurrence of the event of interest is given the return value of the experiment function
// - true: the event of interest ocurred
// - false: the event of intereset did NOT ocurred
func Simulation(experiment func() bool) float64 {
	ocurrencesEvent := 0

	for trial := 0; trial < numTrials; trial++ {
		eventHappend := experiment()
		if eventHappend {
			ocurrencesEvent++
		}
	}

	return float64(ocurrencesEvent) / float64(numTrials)
}
