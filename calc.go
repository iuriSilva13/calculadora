package main

import (
	"fmt"
	"strconv"
)

func main() {
	var primeiroDigito , segundoDigito string
	var operador string
	var resultado int

	fmt.Print("digite um numero:")
	fmt.Scan(&primeiroDigito)

	fmt.Print("digite a operação:")
	fmt.Scan(&operador)

	fmt.Print("digite outro numero:")
	fmt.Scan(&segundoDigito)

	primeiroValor, err := strconv.Atoi(primeiroDigito)
	if err != nil {
		fmt.Println("Erro !", err)
		return
	}

	segundoValor, err := strconv.Atoi(segundoDigito)
	if err != nil {
		fmt.Println("Erro !", err)
		return
	}

	switch operador {
	case "+":
		resultado = primeiroValor + segundoValor
	case "-":
		resultado = primeiroValor - segundoValor
	case "*":
		resultado = primeiroValor * segundoValor
	case "/":
		resultado = primeiroValor / segundoValor
	default:
		fmt.Println("Nenhum operador foi digitado")
		return
	}

	fmt.Println(primeiroValor,operador,segundoValor,"=",resultado)
}
