package main

import "fmt"

var input string

func foo(s string) string {
	return fmt.Sprintf("%s passed as an argument", s)
}

func main() {
	input = "bar"
	value := foo(input)
	fmt.Println(value)

}
