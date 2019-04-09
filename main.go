package main

import "fmt"

type Mooing interface {
	Moo() string
}

type Grazing interface {
	EatGrass()
}

type Cow struct{}

func (c *Cow) Moo() string {
	return "moo"
}

func (c *Cow) EatGrass() {
	fmt.Println("Eating grass")
}

func Milk(cow interface {
	Mooing
	Grazing
}) {
	cow.Moo()
	cow.EatGrass()
}

func main() {
	Milk(&Cow{})
}
