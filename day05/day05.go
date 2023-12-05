package day05

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strings"
	"sync"
)

const filename = "./day05/input.txt"

func strToNumArr(input string) [][3]int {
	strs := strings.Split(input, "\n")
	numArray := make([][3]int, len(strs)-1)
	for i, str := range strs[1:] {
		tmp := utils.MapToInt(strings.Fields(str))
		numArray[i][0] = tmp[0]
		numArray[i][1] = tmp[1]
		numArray[i][2] = tmp[2]
	}
	return numArray
}

var seeds, seedSoil, soilFert, fertWater, waterLight, lightTemp, tempHumid, humidLoc = func() ([]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int) {
	input := strings.Split(utils.ReadString(filename), "\n\n")
	seeds := utils.MapToInt(strings.Split(strings.Split(input[0], ": ")[1], " "))
	return seeds, strToNumArr(input[1]), strToNumArr(input[2]), strToNumArr(input[3]), strToNumArr(input[4]), strToNumArr(input[5]), strToNumArr(input[6]), strToNumArr(input[7])
}()

func getDestOrDefault(input [][3]int, n int) int {
	for _, val := range input {
		dest := val[0]
		source := val[1]
		howMany := val[2]

		if source <= n && n < source+howMany {
			toAdd := n - source
			return dest + toAdd
		}
	}
	return n
}

func part01() int {
	lowest := -1
	for _, seed := range seeds {
		soil := getDestOrDefault(seedSoil, seed)
		fert := getDestOrDefault(soilFert, soil)
		water := getDestOrDefault(fertWater, fert)
		light := getDestOrDefault(waterLight, water)
		temp := getDestOrDefault(lightTemp, light)
		humid := getDestOrDefault(tempHumid, temp)
		loc := getDestOrDefault(humidLoc, humid)
		if lowest == -1 {
			lowest = loc
		} else {
			if loc < lowest {
				lowest = loc
			}
		}
	}
	return lowest
}

func part02() int {
	lowest := -1
	var wg sync.WaitGroup
	for i := 0; i < len(seeds); i += 2 {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
				soil := getDestOrDefault(seedSoil, j)
				fert := getDestOrDefault(soilFert, soil)
				water := getDestOrDefault(fertWater, fert)
				light := getDestOrDefault(waterLight, water)
				temp := getDestOrDefault(lightTemp, light)
				humid := getDestOrDefault(tempHumid, temp)
				loc := getDestOrDefault(humidLoc, humid)
				if lowest == -1 {
					lowest = loc
				} else {
					if loc < lowest {
						lowest = loc
					}
				}
			}
		}()
	}
	wg.Wait()
	return lowest
}

func Main() {
	fmt.Println("Advent of Code 2023, Day 05")
	fmt.Println(part01())
	fmt.Println(part02())
}
