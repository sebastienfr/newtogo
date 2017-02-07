package main

import "fmt"

type gopher struct {
	Swag       bool
	Name       string
	Experience int
	Emails     []string
}

func main() {
	seb := gopher{
		Swag:       true,
		Name:       "Sébastien",
		Experience: 10,
		Emails:     []string{"seb@domain.com"},
	}

	fmt.Printf(seb.Name)
}
