package carspassing

import "fmt"

func main() {
	A := []int{0, 1, 0, 1, 1}
	result := SolutionCarsPassing(A)
	fmt.Println("SolutionCarsPassing")
	fmt.Println(result)
}

func SolutionCarsPassing(A []int) int {
	count := 0
	zeroCount := 0
	for i := 0; i < len(A); i++ {

		if A[i] == 0 {
			zeroCount++

		} else if A[i] == 1 {
			count += zeroCount

		}
	}

	if count >= 0 && count <= 1000000000 {
		return -1
	}

	return count
}
