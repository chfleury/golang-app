package main

import (
	"math/rand"
	"testing"
)

func TestSum(t *testing.T) {
	randomNumber1 := rand.Intn(100)
	randomNumber2 := rand.Intn(100)

	result := Sum(randomNumber1, randomNumber2)

	expected := randomNumber1 + randomNumber2

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
