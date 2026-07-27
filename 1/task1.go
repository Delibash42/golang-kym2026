package main

type Human struct {
	name string
	age  uint
}

func (h Human) aging() {
	println("мне исполнилось", h.age+1)
}

type Action struct {
	Human
	id string
}

func (a Action) backflip() {
	println("я сделал сальто")
}
func main() {
	a := Action{Human: Human{"vasya", 11},
		id: "765869",
	}
	a.aging()
	a.backflip()
}

// to do усложнить
