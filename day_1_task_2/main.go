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
		numberOfZeros := 0
		dial, numberOfZeros = MakeMove(direction, intValue, dial, 0)
		fmt.Printf("Intermediate dial number: %d\n", dial)
		fmt.Printf("Intermediate number of zeros number: %d\n", numberOfZeros)

		result += numberOfZeros
	}

	fmt.Printf("Result: %d\n", result)
}

func MakeMove(direction string, value int, dial int, numberOfZeros int) (int, int) {
	if value == 0 {
		return dial, numberOfZeros
	} else if direction == "R" {
		if dial == MAX {
			return MakeMove(direction, value-1, MIN, numberOfZeros+1)
		} else {
			return MakeMove(direction, value-1, dial+1, numberOfZeros)
		}
	} else {
		if dial == MIN {
			return MakeMove(direction, value-1, MAX, numberOfZeros)
		} else {
			newNumberOfZeros := numberOfZeros
			if dial == 1 {
				newNumberOfZeros += 1
			}
			return MakeMove(direction, value-1, dial-1, newNumberOfZeros)
		}
	}
}
