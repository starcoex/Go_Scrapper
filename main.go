package main

import "fmt"

func canIDrink(age int) bool {
	switch {
	case age < 10:
		return false
	case age > 18:
		return true

	}
	return true
}

func main() {
	fmt.Println(canIDrink(14))
}
