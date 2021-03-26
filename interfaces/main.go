package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

type GreeterByer interface {
	Greeter
	Byer
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s\n", p.Name)
}

func (p Person) Bye() {
	fmt.Printf("Adios soy %s\n", p.Name)
}

func (p Person) String() string {
	return "Hola soy una persona y mi nombre es: " + p.Name
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola soy %s\n", t)
}

func (t Text) Bye() {
	fmt.Printf("Adios soy %s\n", t)
}

// func GreetAll(gs ...Greeter) {
// 	for _, g := range gs {
// 		g.Greet()
// 		//fmt.Printf("Valor: %v, Tipo: %T\n", g, g)
// 	}
// }

// func ByeAll(bs ...Byer) {
// 	for _, b := range bs {
// 		b.Bye()
// 		//fmt.Printf("\t Valor: %v, Tipo: %T\n", g, g)
// 	}
// }

func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}
}

func main() {
	p := Person{Name: "Jose"}
	//t := Text("Daisy")

	//All(p, t)
	fmt.Println(p)
}
