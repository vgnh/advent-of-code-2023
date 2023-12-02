package day01

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const filename = "./day01/input.txt"

var calibrationDoc = utils.ReadLines(filename)

func part01() int {
	sum := 0
	for _, str := range calibrationDoc {
		var num strings.Builder
		runes := []rune(str)
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				num.WriteRune(runes[i])
				break
			}
		}
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				num.WriteRune(runes[i])
				break
			}
		}
		calibrationVal, _ := strconv.Atoi(num.String())
		sum += calibrationVal
	}
	return sum
}

func part02() int {
	numMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	checkStrs := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	sum := 0
	for _, str := range calibrationDoc {
		var num strings.Builder

		lowestIndex := -1
		lowestChk := ""
		highestIndex := -1
		highestChk := ""

		for _, chk := range checkStrs {
			if i := strings.Index(str, chk); i != -1 {
				if lowestIndex == -1 {
					lowestIndex = i
					lowestChk = chk
				} else if i < lowestIndex {
					lowestIndex = i
					lowestChk = chk
				}
			}

			if i := strings.LastIndex(str, chk); i != -1 {
				if highestIndex == -1 {
					highestIndex = i
					highestChk = chk
				} else if i > highestIndex {
					highestIndex = i
					highestChk = chk
				}
			}
		}

		if _, err := strconv.Atoi(lowestChk); err != nil {
			num.WriteString(numMap[lowestChk])
		} else {
			num.WriteString(lowestChk)
		}
		if _, err := strconv.Atoi(highestChk); err != nil {
			num.WriteString(numMap[highestChk])
		} else {
			num.WriteString(highestChk)
		}

		calibrationVal, _ := strconv.Atoi(num.String())
		sum += calibrationVal
	}
	return sum
}

func Main() {
	fmt.Println("Advent of Code 2023, Day 01")
	fmt.Println(part01())
	fmt.Println(part02())
}
