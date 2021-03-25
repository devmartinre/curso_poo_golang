package main

import "fmt"

func main() {
	i := NewInvoice(
		"Mexico",
		"Mexico City",
		NewCustomer("Jose", "Calle 123 #234", "1234567"),
		[]Item{
			NewItem(1, "Curso de Go", 12.34),
			NewItem(1, "Curso de POO con Go", 54.23),
			NewItem(1, "Curso de Testing con Go", 90.00),
		},
	)

	i.SetTotal()
	fmt.Printf("%+v", i)
}
