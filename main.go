package main

import "fmt"

func superAdd(number ...int) int {
	fmt.Println(number)
	total := 0
	// for index, number := range number {
	// 	fmt.Println(index, number)
	// }
	// for i := 0; i <= len(number); i++ {
	// 	total = total + i
	// 	fmt.Println(total)
	// }
	for _, number := range number {
		total += number
	}

	return total

}

func main() {
	result := superAdd(1, 2, 3, 4, 5)
	fmt.Println(result)
}
