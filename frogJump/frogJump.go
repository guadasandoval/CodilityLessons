package frogjump

import "fmt"

func main() {

	result := SolutionFrogJump(10, 85, 30)
	fmt.Println("SolutionFrogJump")
	fmt.Println(result)

}

func SolutionFrogJump(x, y, d int) int {

	var totalJumps int
	initialPosition := x

	for initialPosition < y {
		initialPosition += d
		totalJumps++
	}

	return totalJumps
}
