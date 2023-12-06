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

	parallel := false

	var wg sync.WaitGroup
	for i, f := range []func() (func() int, func() int){day01.Main, day02.Main, day03.Main, day04.Main, day05.Main} {
		i := i
		f := f
		fn := func() {
			part01, part02 := f()
			fmt.Printf("Advent of Code 2023, Day %02d\n%v\n%v\n", i+1, part01(), part02())
		}

		switch parallel {
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
