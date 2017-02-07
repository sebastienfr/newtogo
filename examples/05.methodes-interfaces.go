package main

type Transport interface {
	Move()
}

type Car struct{}

func (c *Car) Move() {}

type Bike struct{}

func (b *Bike) Move() {}

func ReachLocation(t Transport) {
	t.Move()
}
