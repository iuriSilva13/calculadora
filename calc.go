package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"flag"
)

func main() {
	var numeros []string
	var operadores []string
	execução := flag.Bool("e",false,"calcula na linha de comando")
	flag.Bool("i",false,"calcula no modo interativo")
	flag.Bool("help",false,"mostra uma descrição dos comandos")
	flag.Parse()

	for i, _ := range os.Args {
		if i == 0 {
			continue
		}
		if i == 1{
			continue
		}

		if i%2 == 1 {
			operadores = append(operadores, os.Args[i])
		} else {
			numeros = append(numeros, os.Args[i])
		}
	}

	if *execução == true{
		modoExecução(numeros,operadores)
		return
	}

	if len(numeros)-1 != len(operadores) {
		fmt.Println("Você pode usar a calculadora usando os comandos\n-i e -e, para saber mais detalhes sobre estes\ncomandos digite -help.")
		return
	}
}
func modoExecução(numeros,operadores []string)float64{
	resultado := float64(0)
	operador := "+"
	for i, num := range numeros {
		numeros := tratarValor(num, "Calculo")
		resultado = calcularValores(resultado, numeros, operador)
		if len(operadores) > i {
			operador = operadores[i]
		}
	}
	fmt.Println("O resultado é:", resultado)
	return resultado
}
func modoInterativo(){

}
func calcularValoresDoInput(primeiroDigito, operador, segundoDigito string) float64 {
	fmt.Print("Digite o primeiro numero:")
	fmt.Scan(&primeiroDigito)
	fmt.Print("Digite o operador:")
	fmt.Scan(&operador)
	fmt.Print("Digite outro numero:")
	fmt.Scan(&segundoDigito)

	primTratamento := tratarValor(primeiroDigito, "primeiro digito")
	segunTratamento := tratarValor(segundoDigito, "segundo digito")
	resultado := calcularValores(primTratamento, segunTratamento, operador)
	fmt.Println(primTratamento, operador, segunTratamento, "=", resultado)
	return resultado
}
func calcularMaisValores(segundoDigito, operador string, resultadoAnterior float64) float64 {
	var novoCalculo string

	fmt.Print("Deseja fazer um novo calculo?")
	fmt.Scan(&novoCalculo)
	for {
		if novoCalculo == "sim" {
			fmt.Print("Digite o operador:")
			fmt.Scan(&operador)

			fmt.Print("Digite outro numero:")
			fmt.Scan(&segundoDigito)

			segundoValor := tratarValor(segundoDigito, "segundo digito")
			resultado := calcularValores(resultadoAnterior, segundoValor, operador)
			fmt.Println(resultadoAnterior, operador, segundoDigito, "=", resultado)
			return resultado
		} else {
			exibeErro("programa foi encerrado")
		}
	}
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
		exibeErro(digito + " invalido")
	}

	return valorTratado
}
