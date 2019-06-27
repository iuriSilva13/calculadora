package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var primeiroDigito, operador, segundoDigito string
	var resultado float64

	fmt.Print("digite um numero:")
	fmt.Scan(&primeiroDigito)

	fmt.Print("digite a operação:")
	fmt.Scan(&operador)

	fmt.Print("digite outro numero:")
	fmt.Scan(&segundoDigito)

	primeiroValor, err := tratarValor(primeiroDigito)
	if err != nil {
		fmt.Println("Primeiro digito invalido")
		return
	}

	segundoValor, err := tratarValor(segundoDigito)
	if err != nil {
		fmt.Println("Segundo digito invalido")
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

	fmt.Println(primeiroValor, operador, segundoValor, "=", resultado)
}
func tratarValor(valorDigitado string) (float64, error) {
	valorDigitado = strings.Replace(valorDigitado, ",", ".", -1)
	valorTratado, err := strconv.ParseFloat(valorDigitado, 64)
	if err != nil {
		return 0, err
	}

	return valorTratado, err
}
