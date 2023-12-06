package day06

import (
	"advent-of-code-2023/utils"
	"strconv"
	"strings"
)

const filename = "./inputs/day06.txt"

type raceData struct {
	time, dist int
}

var data = func() []raceData {
	doc := utils.ReadLines(filename)
	times := utils.MapToInt(strings.Fields(strings.Split(doc[0], ":")[1]))
	dists := utils.MapToInt(strings.Fields(strings.Split(doc[1], ":")[1]))
	data := make([]raceData, len(times))
	for i := range data {
		data[i] = raceData{
			time: times[i],
			dist: dists[i],
		}
	}
	return data
}()

func part01() int {
	prod := 1
	for _, rd := range data {
		ways := 0
		for i := 0; i <= rd.time; i++ {
			remainingTime := rd.time - i
			distCanCover := i * remainingTime
			if distCanCover > rd.dist {
				ways++
			}
		}
		prod *= ways
	}
	return prod
}

func part02() int {
	var d, t strings.Builder
	for _, rd := range data {
		t.WriteString(strconv.Itoa(rd.time))
		d.WriteString(strconv.Itoa(rd.dist))
	}
	time, _ := strconv.Atoi(t.String())
	dist, _ := strconv.Atoi(d.String())

	ways := 0
	for i := 0; i <= time; i++ {
		remainingTime := time - i
		distCanCover := i * remainingTime
		if distCanCover > dist {
			ways++
		}
	}
	return ways
}

func Main() (int, func() int, func() int) {
	return 6, part01, part02
}
