package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagvar int

	flag.IntVar(&flagvar, "intp", 42, "help message for int flag")

	flag.Parse()

	fmt.Println("flagvar has value ", flagvar)
}
