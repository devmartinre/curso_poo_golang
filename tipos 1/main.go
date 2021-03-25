package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// Declaración de Alias
type myAlias = course

// Definición de Tipo
type newCourse course

type newBool bool

func (b newBool) String() string {
	if b {
		return "VERDADERO"
	} else {
		return "FALSO"
	}
}

func main() {
	//c := newCourse{name: "Go"}
	//c.Print()
	//fmt.Printf("El tipo es: %T\n", c)

	var b newBool = false
	fmt.Println(b.String())
}
