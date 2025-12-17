package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Reading input from path %s\n", os.Args[1])

	data, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("File content:\n%s\n", string(data))

	lines := strings.Split(string(data), "\n")

	result := 0

	for _, line := range lines {
		result += getLargestNumberWithTwo(line)
	}

	fmt.Println(result)
}

func getLargestNumberWithTwo(input string) int {
	result := 0
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			currentResult, err := strconv.Atoi(string(input[i]) + string(input[j]))
			if err != nil {
				fmt.Println("Error converting to int:", err)
				continue
			}
			if currentResult > result {
				result = currentResult
			}
		}
	}
	return result
}
