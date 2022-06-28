package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	result := Fibonacci(4)
	t.Log(result)
}
