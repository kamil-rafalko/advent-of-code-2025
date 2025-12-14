package main

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	if !isIdValid("123456", 2) {
		t.Errorf("isIdValid(\"1234\", 2) = false, want true")
	}

	if !isIdValid("11", 2) {
		t.Errorf("isIdValid(\"11\", 2) = false, want true")
	}

	if isIdValid("11", 1) {
		t.Errorf("isIdValid(\"11\", 1) = true, want false")
	}

	if !isIdValid("222220", 2) {
		t.Errorf("isIdValid(\"222220\", 2) = false, want true")
	}
}
