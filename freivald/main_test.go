package main

import "testing"

func TestTwoOutput(t *testing.T) {
	// Driver Code
	N := 2
	a := [][]int{{1, 1}, {1, 1}}
	b := [][]int{{1, 1}, {1, 1}}
	c := [][]int{{2, 2}, {2, 2}}

	if multiply(a, b, c, N) {
		println("True: C = A * B")
	} else {
		println("False: C != A * B")
	}
}
