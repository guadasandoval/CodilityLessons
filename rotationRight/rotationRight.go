package rotationRight

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5}
	rotationRight(slc, 4)
}

func rotationRight(a []int, d int) []int {
	var newSlice []int
	slcLen := len(a)

	for i := int(0); i < d; i++ {
		newSlice = a[1:slcLen]
		newSlice = append(newSlice, a[0])
		a = newSlice
	}
	fmt.Println(a)
	return a
}