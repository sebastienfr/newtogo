package main

import (
	"errors"
	"fmt"
)

func validate(str string) (string, error) {
	if len(str) > 4 {
		return str[0:4], errors.New("string to long")
	}

	return str, nil
}

func main() {
	world, err := validate("world")
	if err != nil {
		fmt.Println("Error while processing param :", err)
	}
	fmt.Println("Hello ", world)

	world, err = validate("monde")
	fmt.Println("Bonjour ", world)
}
