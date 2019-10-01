package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"io"
)

func main() {
	var primeiroDigito, segundoDigito float64
	var operador string
	var numeros []string
	var operadores []string
	execução := flag.Bool("e", false, "calcula na linha de comando")
	interativo := flag.Bool("i", false, "calcula no modo interativo")
	help := flag.Bool("help", false, "mostra uma descrição dos comandos")
	flag.Parse()

	print := os.Stdout

	for i, _ := range os.Args {
		if i == 0 {
			continue
		}
		if i == 1 {
			continue
		}
		if i == 2 {
			continue
		}

		if len(os.Args)%2 == 1 {
			fmt.Fprintln(print,"Calculo invalido")
			return
		}

		if i%2 == 1 {
			numeros = append(numeros, os.Args[i])
		} else {
			operadores = append(operadores, os.Args[i])
		}
	}

	file := bufio.NewScanner(os.Stdin)

	if *help == true {
		fmt.Fprintln(print,"-i:Entra no modo interativo\n-e =:Você pode fazer o calculo na linha de comando digitando -e =(seu calculo)\n-help:comando de ajuda")
		return
	}

	if *interativo == true {
		modoInterativo(primeiroDigito, segundoDigito, operador,file,print)
		return
	}

	if *execução == true {
		modoExecução(numeros, operadores,print)
		return
	}

	if len(numeros)-1 != len(operadores) {
		fmt.Fprintln(print,"Você pode usar a calculadora usando os comandos\n-i e -e =, para saber mais detalhes sobre estes\ncomandos digite -help.")
		return
	}
}
func modoExecução(numeros, operadores []string,print io.Writer) float64 {
	resultado := float64(0)
	operador := "+"
	var operadorInvalido string
	for i, num := range numeros {
		numeros, err := tratarValor(num, "Calculo",print)
		if err != nil {
			return 0.0
		}
		resultado, operadorInvalido = calcularValores(resultado, numeros, operador,print)
		if operadorInvalido == "Argumento inválido" {
			return 0.0
		}
		if len(operadores) > i {
			operador = operadores[i]
		}
	}
	fmt.Fprintln(print,"O resultado é:", resultado)
	return resultado
}
func lerInputs(file *bufio.Scanner, digito string,print io.Writer) string {
	fmt.Fprint(print,digito)
	if file.Scan() {
		return file.Text()
	}
	return ""
}
func validarEntradas(primeiroDigito, segundoDigito string, primeiraVez bool,print io.Writer) (float64, float64, error) {
	var primeiroTratamento, segundoTratamento float64
	var err error

	if primeiraVez {
		primeiroTratamento, err = tratarValor(primeiroDigito, "primeiro digito",print)
		if err != nil {
			return 0.0, 0.0, err
		}
	}
	return texto
}
func validarEntradas(primeiroDigito,segundoDigito string,primeiraVez bool)(float64,float64,error){
	var primeiroTratamento,segundoTratamento float64
	var err error

	segundoTratamento, err = tratarValor(segundoDigito, "segundo digito",print)
	if err != nil {
		return 0.0, 0.0, err
	}
	return primeiroTratamento, segundoTratamento, err
}
func obterDadosDosInputs(primeiraVez bool,file *bufio.Scanner,print io.Writer) (float64, float64, string, error) {
	var primeiroDigito, segundoDigito, operador string
	var err error

	if primeiraVez {
		primeiroDigito = lerInputs(file, "Digite o primeiro numero:",print)
	}

	operador = lerInputs(file, "Digite o operador:",print)
	segundoDigito = lerInputs(file, "Digite o segundo numero:",print)
	primeiroValorTratado, segundoValorTratado, err := validarEntradas(primeiroDigito, segundoDigito, primeiraVez,print)
	return primeiroValorTratado, segundoValorTratado, operador, err
}
func modoInterativo(primeiroDigito, segundoDigito float64, operador string,file *bufio.Scanner,print io.Writer) (float64, error) {
	var primeiroResultado float64
	var operadorInvalido string
	var err error
	primeiraVez := true
	contador := 0

	for {
		primeiroDigito, segundoDigito, operador, err = obterDadosDosInputs(primeiraVez,file,print)
		if err != nil {
			return 0.0, err
		}
		if primeiraVez {
			primeiroResultado, operadorInvalido = calcularValores(primeiroDigito, segundoDigito, operador,print)
		} else {
			primeiroDigito, operadorInvalido = calcularValores(primeiroResultado, segundoDigito, operador,print)
		}

		if operadorInvalido == "Argumento inválido" {
			return 0.0, err
		}

		if contador == 0 {
			fmt.Fprintln(print,primeiroDigito, operador, segundoDigito, "=", primeiroResultado)
		}
		if contador >= 1 {
			fmt.Fprintln(print,primeiroResultado, operador, segundoDigito, "=", primeiroDigito)
			primeiroResultado = primeiroDigito
		}

		contador = contador + 1
		novoCalculo := lerInputs(file, "Deseja fazer um novo calculo?",print)

		if novoCalculo != "sim" {
			break
		}
		primeiraVez = false
	}
	return primeiroResultado, err
}
func calcularValores(primeiroValor, segundoValor float64, operador string,print io.Writer) (float64, string) {
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
		mensagemErro := exibeErro("Argumento inválido",print)
		return 0.0, mensagemErro
	}
	return resultado, operador
}
func exibeErro(textoErro string,print io.Writer) string {
	fmt.Fprintln(print,"###", textoErro, "###")
	return textoErro
}
func tratarValor(valorDigitado string, digito string,print io.Writer) (float64, error) {
	valorDigitado = strings.Replace(valorDigitado, ",", ".", -1)
	valorTratado, err := strconv.ParseFloat(valorDigitado, 64)
	if err != nil {
		fmt.Fprintln(print,digito + " invalido")
	}

	return valorTratado, err
}
