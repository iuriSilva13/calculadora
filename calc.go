package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"flag"
)

func main() {
	var primeiroDigito, segundoDigito, novoCalculo, operador string
	var numeros []string
	var operadores []string
	execução := flag.Bool("e",false,"calcula na linha de comando")
	interativo := flag.Bool("i",false,"calcula no modo interativo")
	help := flag.Bool("help",false,"mostra uma descrição dos comandos")
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

	if *help == true{
		fmt.Println("-i:Entra no modo interativo\n-e:Você pode fazer o calculo na linha de comando digitando -e (seu calculo)\n-help:comando de ajuda")
		return
	}

	if *interativo == true{
		modoInterativo(primeiroDigito,segundoDigito,novoCalculo,operador)
		return
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
	var operadorInvalido string
	for i, num := range numeros {
		numeros,err := tratarValor(num, "Calculo")
		if err != nil{
			return numeros
		}
		resultado,operadorInvalido = calcularValores(resultado, numeros, operador)
		if operadorInvalido == "Argumento inválido"{
			return resultado
		}
		if len(operadores) > i {
			operador = operadores[i]
		}
	}
	fmt.Println("O resultado é:", resultado)
	return resultado
}
func modoInterativo(primeiroDigito,segundoDigito,novoCalculo,operador string)float64{
	primeiroResultado := calcularValoresDoInput(primeiroDigito, operador, segundoDigito)
	if primeiroResultado == 0{
		return primeiroResultado
	}

	fmt.Print("Deseja fazer um novo calculo?")
	fmt.Scan(&novoCalculo)

	if novoCalculo == "sim" {
		fmt.Print("Digite o operador:")
		fmt.Scan(&operador)

		fmt.Print("Digite outro numero:")
		fmt.Scan(&segundoDigito)

		segundoValor,err := tratarValor(segundoDigito, "segundo digito")
		if err != nil{
			return segundoValor
		}
		segundoResultado,operadorInvalido := calcularValores(primeiroResultado, segundoValor, operador)
		if operadorInvalido == "Argumento inválido"{
			return segundoResultado
		}
		fmt.Println(primeiroResultado, operador, segundoDigito, "=", segundoResultado)
		segundoResultado = calcularMaisValores(primeiroDigito, operador, segundoResultado)
	} else {
		fmt.Println("programa foi encerrado")
		return primeiroResultado
	}
		return primeiroResultado
}
func calcularValores(primeiroValor, segundoValor float64, operador string) (float64,string) {
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
		mensagemErro := exibeErro("Argumento inválido")
		return resultado,mensagemErro
	}
	return resultado,operador
}
func exibeErro(textoErro string) string {
	fmt.Println("###", textoErro, "###")
	return textoErro
}
func tratarValor(valorDigitado string, digito string) (float64,error) {
	valorDigitado = strings.Replace(valorDigitado, ",", ".", -1)
	valorTratado, err := strconv.ParseFloat(valorDigitado, 64)
	if err != nil {
		fmt.Println(digito + " invalido")
	}

	return valorTratado,err
}
