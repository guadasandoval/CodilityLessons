package binarygap

import (
	"fmt"
	"strconv"
)

func main() {
	var numero int
	fmt.Println("Ingrese un número")
	fmt.Scanf("%d\n", &numero)
	fmt.Println("-----------")
	maxCount := Solution(int64(numero))
	fmt.Println("Respuesta: ", maxCount)
}

func Solution(n int64) int {
	var count int
	var maxCount int
	binaryString := strconv.FormatInt(n, 2)

	for _, binaryCharacter := range binaryString {

		if string(binaryCharacter) == "0" {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}

	return maxCount
}
