package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string)(int, string){
return len(name), strings.ToUpper(name)
}
func repeatMe(words ...string){
	fmt.Println(words)
}



func main() {
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)
	fmt.Println((lenAndUpper("Nico")))
	repeatMe("nico", "Kim")
}
