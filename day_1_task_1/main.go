package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX = 99
const MIN = 0

func main() {
	fmt.Printf("First arg %s\n", os.Args[1])

	data, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("File content: %s\n", string(data))

	lines := strings.Split(string(data), "\n")

	fmt.Printf("Lines: %v\n", lines)

	dial := 50
	result := 0

	for _, line := range lines {
		direction := line[:1]
		value := line[1:]

		fmt.Printf("Direction: %s, Value: %s\n", direction, value)

		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting value to int:", err)
			continue
		}
		dial = MakeMove(direction, intValue, dial)
		fmt.Printf("Intermediate dial number: %d\n", dial)

		if dial == MIN {
			result++
		}
	}

	fmt.Printf("Result: %d\n", result)
}

func MakeMove(direction string, value int, dial int) int {
	if direction == "R" {
		if value+dial > MAX {
			return MakeMove(direction, value-(MAX-dial+1), MIN)
		} else {
			return dial + value
		}
	} else {
		if dial-value < MIN {
			return MakeMove(direction, value-(dial-MIN+1), MAX)
		} else {
			return dial - value
		}
	}
}
