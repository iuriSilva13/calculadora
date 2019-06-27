package main

import (
	"fmt"
)

func main() {
	var primeiroNumero , segundoNumero , resultado float64
	var operação string

	fmt.Print("digite um numero:")
	fmt.Scan(&primeiroNumero)

	fmt.Print("digite a operação:")
	fmt.Scan(&operação)

	fmt.Print("digite outro numero:")
	fmt.Scan(&segundoNumero)

	switch operação {
	case "+":
		resultado = primeiroNumero + segundoNumero
	case "-":
		resultado = primeiroNumero - segundoNumero
	case "*":
		resultado = primeiroNumero * segundoNumero
	case "/":
		resultado = primeiroNumero / segundoNumero
	default:
		fmt.Println("Nenhum operador foi digitado")
		return
	}

	fmt.Println(primeiroNumero,operação,segundoNumero,"=",resultado)
}