package missinginteger

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{4, 3, 2}
	result := SolutionMissingInteger(A)
	fmt.Println("SolutionMissingInteger")
	fmt.Println(result)
}

func SolutionMissingInteger(A []int) int {
	smallestInt := 1

	if (len(A) == 0) || A[len(A)-1] <= 0 {
		return smallestInt
	}

	sort.Ints(A)

	if A[0] > 1 {
		return smallestInt
	}

	for i := 0; i < len(A); i++ {
		if A[i] == smallestInt {
			smallestInt++
		}
	}

	return smallestInt
}
