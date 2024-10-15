package main

import "testing"

func TestFibZero(t *testing.T) {
	res := fib(0)
	if res != 0 {
		t.Fatal("must be 0")
	}
}
