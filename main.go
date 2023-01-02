package main

import "fmt"

func canIDrink(age int) bool {
	if koreaAge := age + 2; koreaAge < 18 {
		fmt.Println(koreaAge)
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(canIDrink(14))
}
