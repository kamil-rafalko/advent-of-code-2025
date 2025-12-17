package main

import (
	"testing"
)

func TestGetLargestNumberWithTwo(t *testing.T) {
	if getLargestNumberWithTwo("2736") != 76 {
		t.Errorf("getLargestNumberWithTwo(\"2736\") != 76")
	}
}
