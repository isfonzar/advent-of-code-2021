package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	windowSize = 3
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//////// Solution /////////
	var previousSum int
	var windowSum int
	var totalIncreases int
	
	for i := 0; i < len(input); i++ {
		c, _ := strconv.Atoi(input[i])

		if i < windowSize {
			windowSum += c
			continue
		}

		if previousSum == 0 {
			previousSum = windowSum
		}

		p, _ := strconv.Atoi(input[i-windowSize])

		windowSum += c
		windowSum -= p

		if (windowSum > previousSum) {
			totalIncreases++
		}

		previousSum = windowSum
	}

	fmt.Println(totalIncreases)
}
