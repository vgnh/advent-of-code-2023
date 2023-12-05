package main

import (
	"advent-of-code-2023/day01"
	"advent-of-code-2023/day02"
	"advent-of-code-2023/day03"
	"advent-of-code-2023/day04"
	"advent-of-code-2023/day05"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	runInParallel := false

	var wg sync.WaitGroup
	for _, fun := range []func(){day01.Main, day02.Main, day03.Main, day04.Main, day05.Main} {
		fn := fun

		switch runInParallel {
		case true:
			wg.Add(1)
			go func() {
				fn()
				wg.Done()
			}()
		case false:
			fn()
		}
	}
	wg.Wait()

	fmt.Printf("\nTime elapsed: %v\n", time.Since(start))
}
