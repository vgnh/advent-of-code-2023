package main

import (
	"advent-of-code-2023/day01"
	"advent-of-code-2023/day02"
	"advent-of-code-2023/day03"
	"advent-of-code-2023/day04"
	"advent-of-code-2023/day05"
	"advent-of-code-2023/day06"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	parallel := false

	var wg sync.WaitGroup
	for _, f := range []func() (int, func() int, func() int){day01.Main, day02.Main, day03.Main, day04.Main, day05.Main, day06.Main} {
		f := f
		fn := func() {
			i, part01, part02 := f()
			fmt.Printf("Advent of Code 2023, Day %02d\n%v\n%v\n", i, part01(), part02())
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
