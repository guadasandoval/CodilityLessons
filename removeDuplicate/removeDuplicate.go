package removeduplicate

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6}
	duplicado, index := removeDuplicateInt(slc)
	fmt.Println("el duplicado es:")
	fmt.Println(duplicado)
	fmt.Println("index")
	fmt.Println(index)
}

func removeDuplicateInt(cantidad []int) ([]int, int) {

	var indices []int
	var noDuplicate int

	for _, index := range cantidad {

		if !contains(cantidad, index) {
			indices = append(indices, index)
		}
	}
	return indices, noDuplicate
}

func contains(s []int, e int) bool {
	var count int
	for _, a := range s {
		if a == e {
			count++
		}
		if count > 1 {
			return true
		}
	}
	return false
}
