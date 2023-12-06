package day03

import (
	"advent-of-code-2023/utils"
	"strconv"
	"unicode"
)

const filename = "./inputs/day03.txt"

var schematic = func() [][]rune {
	strs := utils.ReadLines(filename)
	schematic := make([][]rune, len(strs))
	for i, str := range strs {
		schematic[i] = []rune(str)
	}
	return schematic
}()

func part01() int {
	sum := 0
	for i := range schematic {
		prevNum := false
		n := 0
		for j := range schematic[i] {
			if prevNum {
				if unicode.IsDigit(schematic[i][j]) {
					n2, _ := strconv.Atoi(string(schematic[i][j]))
					n = n*10 + n2
					prevNum = true

					if j == len(schematic[i])-1 {
						nLen := countDigits(n)
						if !numNotAdjToSymbol(i, j-nLen+1, nLen, &schematic) {
							sum += n
						}
					}
				} else {
					prevNum = false
					nLen := countDigits(n)
					if !numNotAdjToSymbol(i, j-nLen, nLen, &schematic) {
						sum += n
					}
				}
			} else {
				if unicode.IsDigit(schematic[i][j]) {
					n, _ = strconv.Atoi(string(schematic[i][j]))
					prevNum = true

					if j == len(schematic[i])-1 {
						nLen := countDigits(n)
						if !numNotAdjToSymbol(i, j-nLen+1, nLen, &schematic) {
							sum += n
						}
					}
				}
			}
		}
	}
	return sum
}

func countDigits(a int) int {
	count := 0
	for a != 0 {
		a /= 10
		count++
	}
	return count
}

func numNotAdjToSymbol(i, j, nLen int, schematic *[][]rune) bool {
	r := len(*schematic)
	c := len((*schematic)[0])

	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+nLen; l++ {
			if k < 0 || k > r-1 || l < 0 || l > c-1 {
				continue
			}
			if k == i && l == j {
				l = l + nLen - 1
				continue
			}

			ch := (*schematic)[k][l]
			if unicode.IsDigit(ch) {
				continue
			} else if ch == '.' {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

type point struct {
	x, y int
}

func part02() int {
	sum := 0
	gearNumbersMap := make(map[point][]int)
	for i := range schematic {
		prevNum := false
		n := 0
		for j := range schematic[i] {
			if prevNum {
				if unicode.IsDigit(schematic[i][j]) {
					n2, _ := strconv.Atoi(string(schematic[i][j]))
					n = n*10 + n2
					prevNum = true

					if j == len(schematic[i])-1 {
						nLen := countDigits(n)
						markGears(i, j-nLen+1, n, nLen, &schematic, &gearNumbersMap)
					}
				} else {
					prevNum = false
					nLen := countDigits(n)
					markGears(i, j-nLen, n, nLen, &schematic, &gearNumbersMap)
				}
			} else {
				if unicode.IsDigit(schematic[i][j]) {
					n, _ = strconv.Atoi(string(schematic[i][j]))
					prevNum = true

					if j == len(schematic[i])-1 {
						nLen := countDigits(n)
						markGears(i, j-nLen+1, n, nLen, &schematic, &gearNumbersMap)
					}
				}
			}
		}
	}
	for _, v := range gearNumbersMap {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	return sum
}

func markGears(i, j, n, nLen int, schematic *[][]rune, gearNumbersMap *map[point][]int) {
	r := len(*schematic)
	c := len((*schematic)[0])

	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+nLen; l++ {
			if k < 0 || k > r-1 || l < 0 || l > c-1 {
				continue
			}
			if k == i && l == j {
				l = l + nLen - 1
				continue
			}
			if (*schematic)[k][l] == '*' {
				(*gearNumbersMap)[point{k, l}] = append((*gearNumbersMap)[point{k, l}], n)
			}
		}
	}
}

func Main() (func() int, func() int) {
	return part01, part02
}
