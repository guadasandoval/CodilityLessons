package missingelement

func main() {
	SolutionMissingElement()
}

func SolutionMissingElement() int {
	A := []int{3, 2, 4, 5}
	lenA := len(A)
	var sum int

	for i := 0; i < lenA; i++ {

		sum += A[i]

	}

	n_num_sum := ((lenA + 1) * (lenA + 2)) / 2

	return (n_num_sum - sum)
}
