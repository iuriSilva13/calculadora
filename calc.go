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
		var primeiroDigito,segundoDigito,novoCalculo,operador string
		primeiroResultado := calcularValoresDoInput(primeiroDigito,operador,segundoDigito)

		for{
			fmt.Print("Deseja fazer um novo calculo?")
			fmt.Scan(&novoCalculo)

			if novoCalculo == "sim"{
				fmt.Print("Digite o operador:")
				fmt.Scan(&operador)

				fmt.Print("Digite outro numero:")
				fmt.Scan(&segundoDigito)

				segundoValor := tratarValor(segundoDigito,"segundo digito")
				segundoResultado := calcularValores(primeiroResultado,segundoValor,operador)
				fmt.Println(primeiroResultado,operador,segundoDigito,"=",segundoResultado)
			}else{
				exibeErro("programa foi encerrado")
				return
		}
	}
}

	resultado := float64(0)
	operador := "+"
	for i, num := range numeros {
		numeros := tratarValor(num,"Calculo")
		resultado = calcularValores(resultado, numeros, operador)
		if len(operadores) > i {
			operador = operadores[i]
		}
	}
	fmt.Println("O resultado é:", resultado)

}
func calcularValoresDoInput(primeiroDigito,operador,segundoDigito string)float64{
	fmt.Print("Digite o primeiro numero:")
	fmt.Scan(&primeiroDigito)
	fmt.Print("Digite o operador:")
	fmt.Scan(&operador)
	fmt.Print("Digite outro numero:")
	fmt.Scan(&segundoDigito)

	primTratamento := tratarValor(primeiroDigito,"primeiro digito")
	segunTratamento := tratarValor(segundoDigito,"segundo digito")
	resultado := calcularValores(primTratamento,segunTratamento,operador)
	fmt.Println(primTratamento,operador,segunTratamento,"=",resultado)
	return resultado
}
func calcularValores(primeiroValor, segundoValor float64, operador string) float64 {
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
		exibeErro("Argumento invalido")
	}
	return resultado
}
func exibeErro(textoErro string) {
	fmt.Println("###", textoErro, "###")
	os.Exit(1)
}
func tratarValor(valorDigitado string, digito string) float64 {
	valorDigitado = strings.Replace(valorDigitado, ",", ".", -1)
	valorTratado, err := strconv.ParseFloat(valorDigitado, 64)
	if err != nil {
		exibeErro(digito+ " invalido")
	}

	return valorTratado
}
