package day05

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strings"
)

const filename = "./day05/input.txt"

type almanacMap struct {
	dst, src, rng int
}

func toAlmanacMap(str string) []almanacMap {
	strs := strings.Split(str, "\n")
	mapArray := make([]almanacMap, len(strs)-1)
	for i, str := range strs[1:] {
		tmp := utils.MapToInt(strings.Fields(str))
		mapArray[i] = almanacMap{
			dst: tmp[0],
			src: tmp[1],
			rng: tmp[2],
		}
	}
	return mapArray
}

var seeds, seedSoil, soilFert, fertWater, waterLight, lightTemp, tempHumid, humidLoc = func() ([]int, []almanacMap, []almanacMap, []almanacMap, []almanacMap, []almanacMap, []almanacMap, []almanacMap) {
	input := strings.Split(utils.ReadString(filename), "\n\n")
	seeds := utils.MapToInt(strings.Split(strings.Split(input[0], ": ")[1], " "))
	return seeds, toAlmanacMap(input[1]), toAlmanacMap(input[2]), toAlmanacMap(input[3]), toAlmanacMap(input[4]), toAlmanacMap(input[5]), toAlmanacMap(input[6]), toAlmanacMap(input[7])
}()

func getDestOrDefault(maps []almanacMap, n int) int {
	for _, am := range maps {
		dst := am.dst
		src := am.src
		rng := am.rng

		if src <= n && n < src+rng {
			toAdd := n - src
			return dst + toAdd
		}
	}
	return n
}

func part01() int {
	soil := getDestOrDefault(seedSoil, seeds[0])
	fert := getDestOrDefault(soilFert, soil)
	water := getDestOrDefault(fertWater, fert)
	light := getDestOrDefault(waterLight, water)
	temp := getDestOrDefault(lightTemp, light)
	humid := getDestOrDefault(tempHumid, temp)
	loc := getDestOrDefault(humidLoc, humid)
	lowest := loc

	for _, seed := range seeds[1:] {
		soil = getDestOrDefault(seedSoil, seed)
		fert = getDestOrDefault(soilFert, soil)
		water = getDestOrDefault(fertWater, fert)
		light = getDestOrDefault(waterLight, water)
		temp = getDestOrDefault(lightTemp, light)
		humid = getDestOrDefault(tempHumid, temp)
		loc = getDestOrDefault(humidLoc, humid)
		if loc < lowest {
			lowest = loc
		}
	}
	return lowest
}

type seedRange struct {
	start, end int
}

func getSourceOrDefault(maps []almanacMap, n int) int {
	for _, am := range maps {
		dst := am.dst
		src := am.src
		rng := am.rng

		if dst <= n && n < dst+rng {
			toAdd := n - dst
			return src + toAdd
		}
	}
	return n
}

func part02() int {
	seedRanges := make([]seedRange, len(seeds)/2)
	c := 0
	for i := range seedRanges {
		seedRanges[i] = seedRange{
			start: seeds[c],
			end:   seeds[c] + seeds[c+1] - 1,
		}
		c += 2
	}

	loc := 0
	for {
		humid := getSourceOrDefault(humidLoc, loc)
		temp := getSourceOrDefault(tempHumid, humid)
		light := getSourceOrDefault(lightTemp, temp)
		water := getSourceOrDefault(waterLight, light)
		fert := getSourceOrDefault(fertWater, water)
		soil := getSourceOrDefault(soilFert, fert)
		seed := getSourceOrDefault(seedSoil, soil)
		for _, sr := range seedRanges {
			if seed >= sr.start && seed <= sr.end {
				return loc
			}
		}
		loc++
	}
}

func Main() {
	fmt.Println("Advent of Code 2023, Day 05")
	fmt.Println(part01())
	fmt.Println(part02())
}
