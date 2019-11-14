package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"io"
	"net/http"
	"net/url"
)

func main() {
	var primeiroDigito, segundoDigito float64
	var operador string
	var numeros []string
	var operadores []string
	webServer := flag.Bool("w", false, "calcula no browser")
	execução := flag.Bool("e", false, "calcula na linha de comando")
	interativo := flag.Bool("i", false, "calcula no modo interativo")
	help := flag.Bool("help", false, "mostra uma descrição dos comandos")
	flag.Parse()

	w := os.Stdout

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
			fmt.Fprintln(w,"Calculo invalido")
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
		fmt.Fprintln(w,"-i:Entra no modo interativo\n-e =:Você pode fazer o calculo na linha de comando digitando -e =(seu calculo)\n-w:Você pode fazer o calculo no browser digitando http://localhost:8080/calculadora?v1=(primeiro valor)&operador=(operador)&v2=(segundo valor)\n-help:comando de ajuda")
		return
	}

	if *webServer == true {
		http.HandleFunc("/calculadora", calculadoraWeb)
		http.ListenAndServe(":8080", nil)
		return
	}

	if *interativo == true {
		modoInterativo(primeiroDigito, segundoDigito, operador,file,w)
		return
	}

	if *execução == true {
		modoExecução(numeros, operadores,w)
		return
	}

	if len(numeros)-1 != len(operadores) {
		fmt.Fprintln(w,"Você pode usar a calculadora usando os comandos\n-i,-e = e -w, para saber mais detalhes sobre estes\ncomandos digite -help.")
		return
	}
}
func calculadoraWeb(w http.ResponseWriter, request *http.Request) {
	var numeros,operadores,valor []string
	parametros,err := url.ParseQuery(request.URL.RawQuery)
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	valor = parametros["calculo"]

	for i,_ := range valor{
		if i%2 == 0 {
			numeros = append(numeros,valor[i])
		}else{
			operadores = append(operadores,valor[i])
		}
	}

	modoExecução(numeros, operadores, w)
}
func modoExecução(numeros, operadores []string,w io.Writer) float64 {
	resultado := float64(0)
	operador := "+"
	var operadorInvalido string
	for i, num := range numeros {
		numeros, err := tratarValor(num, "Calculo",w)
		if err != nil {
			return 0.0
		}
		resultado, operadorInvalido = calcularValores(resultado, numeros, operador,w)
		if operadorInvalido == "Argumento inválido" {
			return 0.0
		}
		if len(operadores) > i {
			operador = operadores[i]
		}
	}
	fmt.Fprint(w,"O resultado é: ", resultado)
	return resultado
}
func lerInputs(file *bufio.Scanner, digito string,w io.Writer) string {
	fmt.Fprint(w,digito)
	if file.Scan() {
		return file.Text()
	}
	return ""
}
func validarEntradas(primeiroDigito, segundoDigito string, primeiraVez bool,w io.Writer) (float64, float64, error) {
	var primeiroTratamento, segundoTratamento float64
	var err error

	if primeiraVez {
		primeiroTratamento, err = tratarValor(primeiroDigito, "primeiro digito",w)
		if err != nil {
			return 0.0, 0.0, err
		}
	}

	segundoTratamento, err = tratarValor(segundoDigito, "segundo digito",w)
	if err != nil {
		return 0.0, 0.0, err
	}
	return primeiroTratamento, segundoTratamento, err
}
func obterDadosDosInputs(primeiraVez bool,file *bufio.Scanner,w io.Writer) (float64, float64, string, error) {
	var primeiroDigito, segundoDigito, operador string
	var err error

	if primeiraVez {
		primeiroDigito = lerInputs(file, "Digite o primeiro numero:",w)
	}

	operador = lerInputs(file, "Digite o operador:",w)
	segundoDigito = lerInputs(file, "Digite o segundo numero:",w)
	primeiroValorTratado, segundoValorTratado, err := validarEntradas(primeiroDigito, segundoDigito, primeiraVez,w)
	return primeiroValorTratado, segundoValorTratado, operador, err
}
func modoInterativo(primeiroDigito, segundoDigito float64, operador string,file *bufio.Scanner,w io.Writer) (float64, error) {
	var primeiroResultado float64
	var operadorInvalido string
	var err error
	primeiraVez := true
	contador := 0

	for {
		primeiroDigito, segundoDigito, operador, err = obterDadosDosInputs(primeiraVez,file,w)
		if err != nil {
			return 0.0, err
		}
		if primeiraVez {
			primeiroResultado, operadorInvalido = calcularValores(primeiroDigito, segundoDigito, operador,w)
		} else {
			primeiroDigito, operadorInvalido = calcularValores(primeiroResultado, segundoDigito, operador,w)
		}

		if operadorInvalido == "Argumento inválido" {
			return 0.0, err
		}

		if contador == 0 {
			fmt.Fprintln(w,primeiroDigito, operador, segundoDigito, "=", primeiroResultado)
		}
		if contador >= 1 {
			fmt.Fprintln(w,primeiroResultado, operador, segundoDigito, "=", primeiroDigito)
			primeiroResultado = primeiroDigito
		}

		contador = contador + 1
		novoCalculo := lerInputs(file, "Deseja fazer um novo calculo?",w)

		if novoCalculo != "sim" {
			break
		}
		primeiraVez = false
	}
	return primeiroResultado, err
}
func calcularValores(primeiroValor, segundoValor float64, operador string,w io.Writer) (float64, string) {
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
		mensagemErro := exibeErro("Argumento inválido",w)
		return 0.0, mensagemErro
	}
	return resultado, operador
}
func exibeErro(textoErro string,w io.Writer) string {
	fmt.Fprintln(w,"###", textoErro, "###")
	return textoErro
}
func tratarValor(valorDigitado string, digito string,w io.Writer) (float64, error) {
	valorDigitado = strings.Replace(valorDigitado, ",", ".", -1)
	valorTratado, err := strconv.ParseFloat(valorDigitado, 64)
	if err != nil {
		fmt.Fprintln(w,digito + " invalido")
	}

	return valorTratado, err
}
