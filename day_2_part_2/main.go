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

	fmt.Printf("File content: %s\n", string(data))

	ranges := strings.Split(string(data), ",")

	result := 0

	for _, single_range := range ranges {
		bounds := strings.Split(single_range, "-")

		boundsInt := make([]int, len(bounds))
		for i := range bounds {
			boundsInt[i], err = strconv.Atoi(bounds[i])
			if err != nil {
				fmt.Println("Error converting range bounds to int:", err)
				continue
			}
		}

		fmt.Printf("Processing range: %s\n", single_range)

		for i := boundsInt[0]; i <= boundsInt[1]; i++ {
			iAsString := strconv.Itoa(i)
			numberOfDigits := len(iAsString)

			for j := 1; j <= numberOfDigits/2; j++ {
				if numberOfDigits%j == 0 {
					valid := isIdValid(iAsString, j)
					if !valid {
						fmt.Printf("%s is not a valid number! Part size was %d\n", iAsString, j)
						result += i
						break
					}
				}
			}
		}
	}
	fmt.Printf("Result: %d\n", result)
}

func isIdValid(id string, partSize int) bool {
	numberOfDigits := len(id)
	lastPattern := id[0:partSize]
	isValid := true
	for i := partSize; i <= numberOfDigits-partSize; i += partSize {
		current := id[i : i+partSize]
		//fmt.Printf("Current is %s, last pattern is %s, i is %d\n", current, lastPattern, i)
		if lastPattern == current {
			isValid = false
		} else {
			return true
		}
		lastPattern = current
	}
	return isValid
}
