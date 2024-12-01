package main

import (
	"aoc2024/aoc2024"
	"fmt"
	"os"
)

func main() {
	aoc2024.DayOne(ReadFile("1"))
}

func ReadFile(fileName string) string {
	file, err := os.ReadFile("inputs/" + fileName + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	return string(file)
}
