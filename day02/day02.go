package day02

import (
	"advent-of-code-2023/utils"
	"strconv"
	"strings"
)

const filename = "./inputs/day02.txt"

var games = utils.ReadLines(filename)

const (
	RED_COUNT   = 12
	GREEN_COUNT = 13
	BLUE_COUNT  = 14
	RED_COLOR   = "red"
	GREEN_COLOR = "green"
	BLUE_COLOR  = "blue"
)

func part01() int {
	sum := 0
	for _, game := range games {
		temp := strings.Split(game, ": ")

		id := strings.Split(temp[0], " ")[1]
		sets := strings.Split(temp[1], "; ")

		countsWithinLimit := true
		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, str := range cubes {
				temp2 := strings.Split(str, " ")
				count, _ := strconv.Atoi(temp2[0])
				color := temp2[1]

				switch color {
				case RED_COLOR:
					if count > RED_COUNT {
						countsWithinLimit = false
					}
				case GREEN_COLOR:
					if count > GREEN_COUNT {
						countsWithinLimit = false
					}
				case BLUE_COLOR:
					if count > BLUE_COUNT {
						countsWithinLimit = false
					}
				}

				if !countsWithinLimit {
					break
				}
			}

			if !countsWithinLimit {
				break
			}
		}
		if countsWithinLimit {
			nid, _ := strconv.Atoi(id)
			sum += nid
		}
	}
	return sum
}

func part02() int {
	sum := 0
	for _, game := range games {
		temp := strings.Split(game, ": ")

		sets := strings.Split(temp[1], "; ")
		colorCountMap := make(map[string]int)

		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, str := range cubes {
				temp2 := strings.Split(str, " ")
				count, _ := strconv.Atoi(temp2[0])
				color := temp2[1]

				if colorCountMap[color] < count {
					colorCountMap[color] = count
				}
			}
		}

		power := 1
		for _, v := range colorCountMap {
			power *= v
		}
		sum += power
	}
	return sum
}

func Main() (func() int, func() int) {
	return part01, part02
}
