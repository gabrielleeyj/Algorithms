package main

import (
	"math/rand"
)

// Matrix Multiplication
// Given three matrices, A, B, C, find if C is a product of A and B
// Where C = A * B, vector N
func freivald(a, b, c [][]int, N int) bool {

	random := make([]int, N)
	// Generate random vector
	for i := 0; i < N; i++ {
		random[i] = rand.Int() % N
	}

	// Compute B * random for evaluating
	// A * (B * random) - (C * random)
	bRandom := make([]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			bRandom[i] = bRandom[i] + b[i][j]*random[j]
		}
	}

	// Compute C * random for evaluating
	cRandom := make([]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			cRandom[i] = cRandom[i] + c[i][j]*random[j]
		}
	}

	// Compute A * (B * random)
	aXBRandom := make([]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			aXBRandom[i] = aXBRandom[i] + a[i][j]*bRandom[j]
		}
	}

	// Final Check if value of expression true
	var ans bool
	for i := 0; i < N; i++ {
		if (aXBRandom[i] - cRandom[i]) != 0 {
			ans = false
		}
		ans = true
	}
	return ans
}

func multiply(a, b, c [][]int, k int) bool {
	var ans bool
	for i := 0; i < k; i++ {
		if freivald(a, b, c, k) == false {
			ans = false
		}
		ans = true
	}
	return ans
}

func main() {

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
