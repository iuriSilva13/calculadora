package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	primeiroDigito := os.Args[1]
	operador := os.Args[2]
	segundoDigito := os.Args[3]

	primeiroValor := tratarValor(primeiroDigito, "primeiro")

	segundoValor := tratarValor(segundoDigito, "segundo")

	var resultado float64

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
