package main

import (
	"fmt"
)

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)

		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("Calling g ", i)
		g(i)
	}

	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 1 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
}

func main() {
	f()
	fmt.Println("Returned normally from f.")
}
