package countdiv

import "fmt"

func main() {
	A := 6
	B := 11
	K := 2
	result := SolutionCountDiv(A, B, K)
	fmt.Println("SolutionCountDiv")
	fmt.Println(result)
}

func SolutionCountDiv(A, B, K int) int {
	var count int
	for i := A; i <= B; i++ {
		if i%2 == 0 {
			count++
		}
	}
	return count
}
