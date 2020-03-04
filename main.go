package main

import "fmt"

func main() {
	str := of("Mohit").append(" ").append("Mamoria").afterLast("M").toString()
	fmt.Println(str)
}
