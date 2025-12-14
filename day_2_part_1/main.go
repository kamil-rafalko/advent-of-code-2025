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

			if numberOfDigits%2 == 0 {
				if iAsString[:numberOfDigits/2] == iAsString[numberOfDigits/2:] {
					result += i
				}
			}
		}
	}
	fmt.Printf("Result: %d\n", result)
}
