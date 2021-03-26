package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct{}

func (p Paypal) Pay() {
	fmt.Println("Pagado por PayPal")
}

type Cash struct{}

func (c Cash) Pay() {
	fmt.Println("Pagado por Efectivo")
}

type CreditCard struct{}

func (c CreditCard) Pay() {
	fmt.Println("Pagado por Tarjeta de Crédito")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint

	fmt.Println("Digite un número de método de pago")
	fmt.Println("\t 1: PayPal 2:Efectivo 3:Tarjeta de Crédito")

	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("debe de digitar un método válido")
	}

	if method > 3 {
		panic("debe de digitar un método válido")
	}

	payMethod := Factory(method)

	payMethod.Pay()
}
