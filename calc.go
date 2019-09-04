package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"flag"
)

func main() {
	var primeiroDigito, segundoDigito float64
	var operador string
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
		if i == 2{
			continue
		}

		if len(os.Args)%2 == 1{
			fmt.Println("Calculo invalido")
			return
		}

		if i%2 == 1 {
			numeros = append(numeros, os.Args[i])
		} else {
			operadores = append(operadores, os.Args[i])
		}
	}

	if *help == true{
		fmt.Println("-i:Entra no modo interativo\n-e =:Você pode fazer o calculo na linha de comando digitando -e =(seu calculo)\n-help:comando de ajuda")
		return
	}

	if *interativo == true{
		modoInterativo(primeiroDigito,segundoDigito,operador)
		return
	}

	if *execução == true{
		modoExecução(numeros,operadores)
		return
	}

	if len(numeros)-1 != len(operadores) {
		fmt.Println("Você pode usar a calculadora usando os comandos\n-i e -e =, para saber mais detalhes sobre estes\ncomandos digite -help.")
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
func obterDadosDosInputs(primeiraVez bool) (float64, float64, string, error){
	var primeiroDigito, segundoDigito, operador string

	if primeiraVez {
		fmt.Print("Digite o primeiro numero:")
		fmt.Scan(&primeiroDigito)
		fmt.Print("Digite o operador:")
		fmt.Scan(&operador)
		fmt.Print("Digite outro numero:")
		fmt.Scan(&segundoDigito)

		primeiroTratamento, err := tratarValor(primeiroDigito, "primeiro digito")
	 	if err != nil {
	 		return 0.0, 0.0, operador,err
	 	}

	 	segundoTratamento, err := tratarValor(segundoDigito, "segundo digito")
	 	if err != nil {
	 		return 0.0, 0.0, operador,err
	 	}
	 	return primeiroTratamento, segundoTratamento, operador,err
	}

	fmt.Print("Digite o operador:")
	fmt.Scan(&operador)
	fmt.Print("Digite outro numero:")
	fmt.Scan(&segundoDigito)

	segundoTratamento, err := tratarValor(segundoDigito, "segundo digito")
	if err != nil {
	 	return 0.0, 0.0, operador,err
	}
  	return 0.0, segundoTratamento, operador,err
}
func modoInterativo(primeiroDigito,segundoDigito float64,operador string)(float64,error){
	var primeiroResultado float64
	var operadorInvalido,novoCalculo string
	var err error
	primeiraVez := true
	contador := 0

	for{
		primeiroDigito, segundoDigito, operador,err = obterDadosDosInputs(primeiraVez)
		if err != nil{
			return 0.0,err
		}
		if primeiraVez {
			primeiroResultado, operadorInvalido = calcularValores(primeiroDigito, segundoDigito, operador)
		}else{
			primeiroDigito, operadorInvalido = calcularValores(primeiroResultado, segundoDigito, operador)
		}

		if operadorInvalido == "Argumento inválido" {
			return 0.0,err
		}

		if contador == 0 {
			fmt.Println(primeiroDigito, operador, segundoDigito, "=", primeiroResultado)
		}
		if contador >= 1 {
		 	fmt.Println(primeiroResultado, operador, segundoDigito, "=", primeiroDigito)
		 	primeiroResultado = primeiroDigito
		}

		contador = contador + 1

		fmt.Print("Deseja fazer um novo calculo?")
		fmt.Scan(&novoCalculo)

		if novoCalculo != "sim" {
			break
		}
		primeiraVez = false
	}
	return primeiroResultado,err
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
