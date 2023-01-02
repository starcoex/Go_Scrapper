package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string)(length int,uppercase string){
	defer fmt.Println("I'm done")
length = len(name)
uppercase= strings.ToLower(name)
return
}
func repeatMe(words ...string){
	fmt.Println(words)
}



func main() {
totalLen, up := lenAndUpper("kKKEK")
fmt.Println(totalLen, up)
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)
	fmt.Println((lenAndUpper("Nico")))
	repeatMe("nico", "Kim")
}
