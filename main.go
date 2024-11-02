package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func sumOfDivisors(n int) int {
	divisorsSum := 1
	sqrtN := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			divisorsSum += i
			if i != n/i {
				divisorsSum += n / i
			}
		}
	}
	return divisorsSum
}

func findAmicableNumbersInRange(start, end int, wg *sync.WaitGroup, mu *sync.Mutex, result *[][2]int) {
	defer wg.Done()
	localResults := [][2]int{}

	for num := start; num < end; num++ {
		sumDiv := sumOfDivisors(num)
		if sumDiv > num && sumOfDivisors(sumDiv) == num {
			localResults = append(localResults, [2]int{num, sumDiv})
		}
	}

	mu.Lock()
	*result = append(*result, localResults...)
	mu.Unlock()
}

func parallelFindAmicableNumbers(limit int) [][2]int {
	numProcesses := runtime.NumCPU()
	chunkSize := limit / numProcesses
	var wg sync.WaitGroup
	var mu sync.Mutex
	result := [][2]int{}

	for i := 0; i < numProcesses; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numProcesses-1 {
			end = limit
		}
		wg.Add(1)
		go findAmicableNumbersInRange(start, end, &wg, &mu, &result)
	}

	wg.Wait()
	return result
}

func main() {
	limit := 1000000 // Imposta il numero che vuoi testare
	startTime := time.Now()

	amicablePairs := parallelFindAmicableNumbers(limit)
	endTime := time.Now()

	fmt.Printf("Coppie di numeri amici trovate fino a %d:\n", limit)
	for _, pair := range amicablePairs {
		fmt.Println(pair)
	}

	fmt.Printf("\nTempo impiegato: %v secondi\n", endTime.Sub(startTime).Seconds())
}
