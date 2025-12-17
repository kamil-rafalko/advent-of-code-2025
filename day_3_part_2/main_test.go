package main

import (
	"testing"
)

func TestGetLargestNumberWithTwelveRecursive(t *testing.T) {
	result := getLargestNumberWithTwelveRecursive("9876543211111")
	if result != 987654321111 {
		t.Errorf("getLargestNumberWithTwo(\"987654321111111\") [%d] != 987654321111", result)
	}

	result = getLargestNumberWithTwelveRecursive("811111111111119")
	if result != 811111111119 {
		t.Errorf("getLargestNumberWithTwo(\"811111111111119\") [%d] != 811111111119", result)
	}

	result = getLargestNumberWithTwelveRecursive("234234234234278")
	if result != 434234234278 {
		t.Errorf("getLargestNumberWithTwo(\"234234234234278\") [%d] != 434234234278", result)
	}

	result = getLargestNumberWithTwelveRecursive("818181911112111")
	if result != 888911112111 {
		t.Errorf("getLargestNumberWithTwo(\"818181911112111\") [%d] != 888911112111", result)
	}
}

func TestGetLargestNumberWithTwelve(t *testing.T) {
	result := getLargestNumberWithTwelve([]int{}, "9876543211111")
	if result != 987654321111 {
		t.Errorf("getLargestNumberWithTwo(\"987654321111111\") [%d] != 987654321111", result)
	}

	result = getLargestNumberWithTwelve([]int{}, "811111111111119")
	if result != 811111111119 {
		t.Errorf("getLargestNumberWithTwo(\"811111111111119\") [%d] != 811111111119", result)
	}

	result = getLargestNumberWithTwelve([]int{}, "234234234234278")
	if result != 434234234278 {
		t.Errorf("getLargestNumberWithTwo(\"234234234234278\") [%d] != 434234234278", result)
	}

	result = getLargestNumberWithTwelve([]int{}, "818181911112111")
	if result != 888911112111 {
		t.Errorf("getLargestNumberWithTwo(\"818181911112111\") [%d] != 888911112111", result)
	}
}

func TestGetLargestNumberWithTwelveStack(t *testing.T) {
	result := getLargestNumberWithTwelveStack("9876543211111")
	if result != 987654321111 {
		t.Errorf("getLargestNumberWithTwo(\"987654321111111\") [%d] != 987654321111", result)
	}

	result = getLargestNumberWithTwelveStack("811111111111119")
	if result != 811111111119 {
		t.Errorf("getLargestNumberWithTwo(\"811111111111119\") [%d] != 811111111119", result)
	}

	result = getLargestNumberWithTwelveStack("234234234234278")
	if result != 434234234278 {
		t.Errorf("getLargestNumberWithTwo(\"234234234234278\") [%d] != 434234234278", result)
	}

	result = getLargestNumberWithTwelveStack("818181911112111")
	if result != 888911112111 {
		t.Errorf("getLargestNumberWithTwo(\"818181911112111\") [%d] != 888911112111", result)
	}
}
