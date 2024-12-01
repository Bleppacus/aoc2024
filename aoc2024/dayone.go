package aoc2024

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func DayOne(file string) {
	var arr []int
	var arr2 []int
	result := strings.Split(file, "\n")
	for i := 0; i < len(result); i++ {
		line := strings.Split(result[i], "   ")
		arr = append(arr, ParseToInt(line[0]))
		arr2 = append(arr2, ParseToInt(line[1]))
	}
	sort.Ints(arr)
	sort.Ints(arr2)
	GetTotalDistance(arr, arr2)
	GetTotalSimilarity(arr, arr2)
}
func ParseToInt(input string) int {
	parsed, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return parsed
}

func GetTotalDistance(input1 []int, input2 []int) int {
	var totalDistance int
	for i := 0; i < len(input1); i++ {
		sub := input1[i] - input2[i]
		if sub < 0 {
			sub = sub * -1
		}
		totalDistance = totalDistance + sub
	}
	fmt.Println("totalDistance:", totalDistance)
	return totalDistance
}
func GetTotalSimilarity(input1 []int, input2 []int) int {
	var totalSimilarity int
	for i := 0; i < len(input1); i++ {
		array1IndexValue := input1[i]
		var sim int
		for j := 0; j < len(input2); j++ {
			if array1IndexValue == input2[j] {
				sim++
			}
		}
		totalSimilarity = totalSimilarity + (array1IndexValue * sim)
	}
	fmt.Println("totalSimilarity:", totalSimilarity)
	return totalSimilarity
}
