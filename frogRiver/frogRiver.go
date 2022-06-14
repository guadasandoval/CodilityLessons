package frogriver

import "fmt"

func main() {

	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	result := SolutionFrogRiverOne(5, A)
	fmt.Println("SolutionFrogRiverOne")
	fmt.Println(result)
}

func SolutionFrogRiverOne(X int, A []int) int {

	for i := 0; i < len(A); i++ {
		if A[i] == X {
			return i
		}
	}
	return -1
}
