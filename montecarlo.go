// Package montecarlo provides a simple interface to execute Monte Carlo simulations.
package montecarlo

import (
	"sync"
)

const numTrials = 2_000_000

// Returns the probability of the event of interest after running simulations
//
// The occurrence of the event of interest is given the return value of the experiment function
// - true: the event of interest ocurred
// - false: the event of intereset did NOT ocurr
func Simulation(experiment func() bool) float64 {
	return simulationSequential(experiment)
}

func simulationSequential(experiment func() bool) float64 {
	ocurrencesEvent := 0

	for trial := 0; trial < numTrials; trial++ {
		eventHappend := experiment()
		if eventHappend {
			ocurrencesEvent++
		}
	}

	return float64(ocurrencesEvent) / float64(numTrials)
}

const nGoroutines = 1

func simulationConcurrent(experiment func() bool) float64 {
	ch := make(chan int)
	var wg sync.WaitGroup

	// Launch work in multiple goroutines
	for i := 0; i < nGoroutines; i++ {
		wg.Add(1)
		go func() {
			localOcurrences := 0
			for j := 0; j < numTrials/nGoroutines; j++ {
				eventHappend := experiment()
				if eventHappend {
					localOcurrences++
				}
			}
			ch <- localOcurrences
			wg.Done()
		}()
	}

	// Close the channel when all the goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Acummulate the results of each experiment
	ocurrencesEvent := 0
	for localOcurrences := range ch {
		ocurrencesEvent += localOcurrences
	}

	return float64(ocurrencesEvent) / float64(numTrials)
}
