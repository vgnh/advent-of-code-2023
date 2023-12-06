package day04

import (
	"advent-of-code-2023/utils"
	"strconv"
	"strings"
)

const filename = "./day04/input.txt"

var scratchCards = utils.ReadLines(filename)

func part01() int {
	sum := 0
	for _, card := range scratchCards {
		points := 0
		strs := strings.Split(strings.Split(card, ": ")[1], " | ")
		winningNumbers := utils.MapToInt(strings.Fields(strs[0]))
		ourNumbers := utils.MapToInt(strings.Fields(strs[1]))
		winningMap := make(map[int]struct{})
		for _, n := range winningNumbers {
			winningMap[n] = struct{}{}
		}
		for _, n := range ourNumbers {
			if _, ok := winningMap[n]; ok {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		sum += points
	}
	return sum
}

func part02() int {
	processed := 0

	originalMap := make(map[int]string)
	afterProcessMap := make(map[int][]int)
	processQueue := []int{}
	for _, card := range scratchCards {
		tmp := strings.Split(card, ": ")
		id, _ := strconv.Atoi(strings.Fields(tmp[0])[1])
		originalMap[id] = tmp[1]

		cardsAfterProcess := processCard(id, &originalMap)
		afterProcessMap[id] = cardsAfterProcess
		processQueue = append(processQueue, cardsAfterProcess...)

		processed++
	}

	for len(processQueue) > 0 {
		card := processQueue[0]
		processQueue = processQueue[1:]

		if v, ok := afterProcessMap[card]; ok {
			if len(v) > 0 {
				processQueue = append(processQueue, v...)
			}
		} else {
			processQueue = append(processQueue, processCard(card, &originalMap)...)
		}

		processed++
	}

	return processed
}

func processCard(id int, originalMap *map[int]string) []int {
	cardsAfterProcess := []int{}

	strs := strings.Split((*originalMap)[id], " | ")
	winningNumbers := utils.MapToInt(strings.Fields(strs[0]))
	ourNumbers := utils.MapToInt(strings.Fields(strs[1]))

	winningMap := make(map[int]struct{})
	for _, n := range winningNumbers {
		winningMap[n] = struct{}{}
	}
	matches := 0
	for _, n := range ourNumbers {
		if _, ok := winningMap[n]; ok {
			matches++
		}
	}

	inc := 0
	for matches > 0 {
		inc++
		cardsAfterProcess = append(cardsAfterProcess, id+inc)
		matches--
	}

	return cardsAfterProcess
}

func Main() (func() int, func() int) {
	return part01, part02
}
