package main

import (
	"fmt"
)

var (
	myVarOne = "ONE"
	myVarTwo = "TWO"
)

var globalURL = "http://localhost:3000"

func main() {
	fmt.Println("Golang Variables:")

	var i int
	i = 10

	var j string = "Pradeep"

	k := 34

	globalURL = "http://localhost:3001"

	fmt.Printf("%v \n", i)
	fmt.Printf("%v \n", j)
	fmt.Printf("%v \n", k)
	fmt.Printf("%v \n", globalURL)
	fmt.Printf("%v \n", myVarOne)
	fmt.Printf("%v \n", myVarTwo)

}
