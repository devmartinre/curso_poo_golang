package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	X()
}

func wrapper(i interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", i, i)

	// v, ok := i.(string)
	// if ok {
	// 	fmt.Printf("\t%s\n", strings.ToUpper(v))
	// }

	switch v := i.(type) {
	case string:
		fmt.Printf("\t%s\n", strings.ToUpper(v))
	case int:
		fmt.Println(v * 2)
	default:
		fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
	}
}

func main() {
	// var e exampler
	// fmt.Printf("Valor: %v, Tipo: %T", e, e)
	wrapper("Martin")
	wrapper(12)
	wrapper(14.32)
	wrapper(false)
	wrapper("Jose")
}
