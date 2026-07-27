package main

import "fmt"

type Animal interface {
	Voice() string
}

type Cat struct {
	Name string
}

func (c Cat) Voice() string {
	return "meow"
}

type Dog struct {
	Name string
}

func (d Dog) Voice() string {
	return "bark"
}

type Cow struct {
	Name string
}

func (c Cow) Voice() string {
	return "moo"
}
func chorus(animal []Animal) {
	for _, v := range animal {
		fmt.Println(v.Voice())
	}
}
func main() {
	animalsList := []Animal{
		Dog{Name: "Шарик"},
		Cat{Name: "Мурка"},
		Cow{Name: "Буренка"},
	}
	chorus(animalsList)
}
