package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	strs := strings.Split(strings.TrimSpace(string(data)), "\n")
	for i := range strs {
		strs[i] = strings.TrimSpace(strs[i])
	}
	return strs
	/* file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines */
}

func ReadString(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	return strings.TrimSpace(string(data))
}

func MapToInt(strs []string) []int {
	ints := make([]int, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err.Error())
		}
		ints[i] = num
	}
	return ints
}

func MapToStr(ints []int) []string {
	strs := make([]string, len(ints))
	for i, num := range ints {
		strs[i] = strconv.Itoa(num)
	}
	return strs
}

func Sum[T int | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func Count[T comparable](slice []T, item T) int {
	count := 0
	for _, v := range slice {
		if v == item {
			count++
		}
	}
	return count
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

func Copy2d[T any](matrix [][]T) [][]T {
	dup := make([][]T, len(matrix))
	for i := range matrix {
		dup[i] = make([]T, len(matrix[i]))
		copy(dup[i], matrix[i])
	}
	return dup
}
