package permutation

import "fmt"

func main() {
	A := []int{4, 1, 3, 2}
	result := SolutionPermutation(A)
	fmt.Println("SolutionPermutation")
	fmt.Println(result)
}
func SolutionPermutation(A []int) int {

	var maxInt int
	for _, element := range A {
		if element > maxInt {
			maxInt = element
		}
	}

	if maxInt != len(A) {
		return 0
	}
	return 1
}
