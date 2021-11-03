package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println("Hello World")
	fmt.Println("My name is Treveon")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(add(33, 36), ",nice")
}
