package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numeros []string
	var operadores []string

	for i, _ := range os.Args {
		if i == 0 {
			continue
		}

		if i%2 == 1 {
			numeros = append(numeros, os.Args[i])
		} else {
			operadores = append(operadores, os.Args[i])
		}
	}

	if len(numeros)-1 != len(operadores) {
		var primeiroDigito,operador,segundoDigito string

		fmt.Print("Digite o primeiro numero:")
		fmt.Scan(&primeiroDigito)
		fmt.Print("Digite o operador:")
		fmt.Scan(&operador)
		fmt.Print("Digite outro numero:")
		fmt.Scan(&segundoDigito)

		tratarPrimeiroDigito := tratarValor(primeiroDigito,"primeiro digito")
		tratarSegundoDigito := tratarValor(segundoDigito,"segundo digito")
		fmt.Println(tratarPrimeiroDigito,operador,tratarSegundoDigito)
		return
	}

	resultado := float64(0)
	operador := "+"
	for i, num := range numeros {
		numeros := tratarValor(num, i)
		resultado = calcularValores(resultado, numeros, operador, i)
		if len(operadores) > i {
			operador = operadores[i]
		}
	}
	fmt.Println(resultado)

}
func calcularValores(primeiroValor, segundoValor float64, operador string, digito int) float64 {
	var resultado float64
	switch operador {
	case "+":
		resultado = primeiroValor + segundoValor
	case "-":
		resultado = primeiroValor - segundoValor
	case "/":
		resultado = primeiroValor / segundoValor
	case "*":
		resultado = primeiroValor * segundoValor
	default:
		exibeErro(fmt.Sprintf("O argumento %d deve ser um operador. Passado: %s", digito, operador))
	}
	return resultado
}
func exibeErro(textoErro string) {
	fmt.Println("###", textoErro, "###")
	os.Exit(1)
}
func tratarValor(valorDigitado string, digito int) float64 {
	valorDigitado = strings.Replace(valorDigitado, ",", ".", -1)
	valorTratado, err := strconv.ParseFloat(valorDigitado, 64)
	if err != nil {
		exibeErro(fmt.Sprintf("O argumento %d deve ser um número", digito))
	}

	return valorTratado
}
