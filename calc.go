package main

import (
	"fmt"
	"os"
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

	primeiroValor := tratarValor(primeiroDigito, "primeiro")

	segundoValor := tratarValor(segundoDigito, "segundo")

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
func exibeErro(textoErro string) {
	fmt.Println("###", textoErro, "###")
	os.Exit(1)
}
func tratarValor(valorDigitado string, digito string) float64 {
	valorDigitado = strings.Replace(valorDigitado, ",", ".", -1)
	valorTratado, err := strconv.ParseFloat(valorDigitado, 64)
	if err != nil {
		exibeErro(digito + " digito é invalido")
	}

	return valorTratado
}
