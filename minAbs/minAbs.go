package minabs

import "math"

func main() {
	A := []int{3, 1, 2, 4, 3}
	SolutionMinAbs(A)
}

func SolutionMinAbs(A []int) int {
	lenA := len(A)

	var sumA int
	for i := 0; i < lenA; i++ {

		sumA += A[i]

	}

	var sumP int
	r := 100000
	for p := 1; p < lenA; p++ {
		sumP += A[p-1]

		s2 := sumA - sumP

		t := math.Abs(float64(sumP - s2))

		if int(t) < r {
			r = int(t)
		}
	}

	return r
}
