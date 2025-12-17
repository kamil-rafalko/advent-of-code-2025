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
		result += getLargestNumberWithTwelve([]int{}, line)
	}

	fmt.Println(result)
}

// To slow
func getLargestNumberWithTwelveRecursive(input string) int {
	if len(input) == 12 {
		result, _ := strconv.Atoi(input)
		return result
	} else {
		largest := 0
		for i, _ := range input {
			current := getLargestNumberWithTwelveRecursive(input[:i] + input[i+1:])
			if current > largest {
				largest = current
			}
		}
		return largest
	}
}

func getLargestNumberWithTwelveStack(input string) int {
	var digits []int
	for i, digit := range input {
		digitConverted, _ := strconv.Atoi(string(digit))
		if len(digits) == 0 {
			digits = append(digits, digitConverted)
		} else {
			minIndex := 0
			diff := 12 - len(input) + i
			if diff > 0 {
				minIndex = diff
			}
			maxIndex := 12
			if len(digits) < 12 {
				maxIndex = len(digits)
			}
			inserted := false
			for j := minIndex; j < maxIndex; j++ {
				if digits[j] < digitConverted {
					digits[j] = digitConverted
					digits = digits[:j+1]
					inserted = true
					break
				}
			}
			if inserted == false && len(digits) < 12 {
				digits = append(digits, digitConverted)
			}
		}
	}
	var stringBuilder strings.Builder
	for _, digit := range digits {
		stringBuilder.WriteString(strconv.Itoa(digit))
	}
	result, _ := strconv.Atoi(stringBuilder.String())
	return result
}

func getLargestNumberWithTwelve(indexes []int, input string) int {
	if len(indexes) == 12 {
		var stringBuilder strings.Builder

		for _, i := range indexes {
			stringBuilder.WriteString(string(input[i]))
		}
		result, _ := strconv.Atoi(stringBuilder.String())
		return result
	} else {
		var largests []int
		largest := 0
		startIndex := 0
		if len(indexes) > 0 {
			startIndex = indexes[len(indexes)-1] + 1
		}
		for i := startIndex; i <= len(input)-12+len(indexes); i++ {
			current, _ := strconv.Atoi(string(input[i]))
			if current > largest {
				largest = current
				largests = []int{i}
			} else if current == largest {
				largests = append(largests, i)
			}
		}

		largestNumber := 0
		for _, index := range largests {
			currentNumber := getLargestNumberWithTwelve(append(indexes, index), input)
			if largestNumber < currentNumber {
				largestNumber = currentNumber
			}
		}
		return largestNumber
	}
}
