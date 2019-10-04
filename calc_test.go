package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	os.Mkdir("testes",0777)
	output := m.Run()
	os.RemoveAll("testes")
	os.Exit(output)
}
func Test_calcularValores(teste *testing.T) {
	type parâmetrosRecebidos struct {
		primeiroValor float64
		segundoValor  float64
		operador      string
	}

	casosDeTestes,err := os.Create("./testes/casos de teste da funcao calcularValores.txt")
	if err != nil{
		teste.Fatal(err)
	}

	testes := []struct {
		mensagemDeIdentificação string
		parâmetrosRecebidos     func(teste *testing.T) parâmetrosRecebidos

		print         *os.File
		valorEsperado float64
		erroEsperado  string
	}{
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroValor: -10,
					segundoValor:  -5,
					operador:      "+",
				}
			},
			valorEsperado: -15.0,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Float com . deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroValor: -80.3,
					segundoValor:  -4.2,
					operador:      "/",
				}
			},
			valorEsperado: 19.119047619047617,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Operador de subtração deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroValor: 45,
					segundoValor:  35,
					operador:      "-",
				}
			},
			valorEsperado: 10.0,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Operador de multiplicação deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroValor: 10,
					segundoValor:  5,
					operador:      "*",
				}
			},
			valorEsperado: 50.0,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Operador deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroValor: 10,
					segundoValor:  45,
					operador:      "gfdgfdgfd",
				}
			},
			valorEsperado: 0.0,
			erroEsperado:  "Argumento inválido",
			print: casosDeTestes,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			testeCalcularValores := valorTeste.parâmetrosRecebidos(teste)

			valorRecebido, err := calcularValores(testeCalcularValores.primeiroValor, testeCalcularValores.segundoValor, testeCalcularValores.operador, valorTeste.print)

			if !reflect.DeepEqual(valorRecebido, valorTeste.valorEsperado) {
				teste.Errorf("calcularValores erro = %v, valorRecebido = %v, valorEsperado: %v", err, valorRecebido, valorTeste.valorEsperado)
			}
		})
	}
}
func Test_validarEntradas(teste *testing.T) {
	type parâmetrosRecebidos struct {
		primeiroDigito string
		segundoDigito  string
		primeiraVez    bool
		erro           error
		print          *os.File
	}

	casosDeTestes,err := os.Create("./testes/casos de teste da funcao validarEntradas.txt")
	if err != nil{
		teste.Fatal(err)
	}

	testes := []struct {
		mensagemDeIdentificação string
		parâmetrosRecebidos     func(teste *testing.T) parâmetrosRecebidos

		print                 *os.File
		primeiroValorEsperado float64
		segundoValorEsperado  float64
		primeiraVez           bool
		erroEsperado          error
	}{
		{
			mensagemDeIdentificação: "caso de erro com valor false deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroDigito: "",
					segundoDigito:  "",
					primeiraVez:    false,
				}
			},
			primeiroValorEsperado: 0.0,
			segundoValorEsperado:  0.0,
			primeiraVez:		   false,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroDigito: "4",
					segundoDigito:  "5",
					primeiraVez:    true,
					erro:           nil,
				}
			},
			primeiroValorEsperado: 4.0,
			segundoValorEsperado:  5.0,
			primeiraVez:           true,
			erroEsperado:          nil,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroDigito: "4",
					segundoDigito:  "5",
					primeiraVez:    false,
					erro:           nil,
				}
			},
			primeiroValorEsperado: 4.0,
			segundoValorEsperado:  5.0,
			primeiraVez:           false,
			erroEsperado:          nil,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "caso de erro com valor true deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroDigito: "",
					segundoDigito:  "",
					primeiraVez:    true,
				}
			},
			primeiroValorEsperado: 0.0,
			segundoValorEsperado:  0.0,
			primeiraVez:           true,
			print: casosDeTestes,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			testeValidarEntradas := valorTeste.parâmetrosRecebidos(teste)

			primeiroValor, segundoValor, err := validarEntradas(testeValidarEntradas.primeiroDigito, testeValidarEntradas.segundoDigito, testeValidarEntradas.primeiraVez, testeValidarEntradas.print)

			if !reflect.DeepEqual(segundoValor, valorTeste.segundoValorEsperado) {
				teste.Errorf("validarEntradas primeiroValorRecebido = %v,segundoValorRecebido = %v,primeiroCalculoRecebido = %v,primeiroValorEsperado = %v,segundoValorEsperado = %v, erroEsperado = %v", primeiroValor, segundoValor, err, valorTeste.primeiroValorEsperado, valorTeste.segundoValorEsperado, valorTeste.erroEsperado)
			}
		})
	}
}
func Test_obterDadosDosInputs(teste *testing.T) {
	casosDeTestes,err := os.Create("./testes/casos de teste da funcao obterDadosDosInputs.txt")
	if err != nil{
		teste.Fatal(err)
	}

	testes := []struct {
		mensagemDeIdentificação string
		primeiraVez             bool
		primeiroDigito          float64
		segundoDigito           float64
		operador                string
		err                     error
		input                   string
		print                   *os.File
	}{
		{
			mensagemDeIdentificação: "Float64 deve ser identificado corretamente",
			primeiraVez:             false,
			primeiroDigito:          25.0,
			segundoDigito:           4.0,
			operador:                "+",
			err:                     nil,
			input:                   "25.0\n+\n4.0\n",
			print: casosDeTestes,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			file, err := ioutil.TempFile("", "")
			if err != nil {
				teste.Fatal(err)
			}

			defer file.Close()

			_, err = io.WriteString(file, valorTeste.input)
			if err != nil {
				teste.Fatal(err)
			}

			_, err = file.Seek(0, os.SEEK_SET)
			if err != nil {
				teste.Fatal(err)
			}
			input := bufio.NewScanner(file)

			primeiroDigito, segundoDigito, operador, err := obterDadosDosInputs(valorTeste.primeiraVez, input, valorTeste.print)

			if valorTeste.primeiraVez != false {
				teste.Errorf("primeiro digito recebido = %v ,segundo digito recebido = %v,operador recebido = %v, primeiro digito esperado = %v, segundo digito esperado = %v,operador esperado = %v ", primeiroDigito, segundoDigito, operador, valorTeste.primeiroDigito, valorTeste.segundoDigito, valorTeste.operador)
			}
		})
	}
}
func Test_modoInterativo(teste *testing.T) {
	casosDeTestes,err := os.Create("./testes/casos de teste da funcao modoInterativo.txt")
	if err != nil{
		teste.Fatal(err)
	}

	testes := []struct {
		mensagemDeIdentificação string
		resultado               float64
		primeiroDigito          float64
		segundoDigito           float64
		operador                string
		input                   string
		print                   *os.File
	}{
		{
			mensagemDeIdentificação: "Float64 deve ser identificado corretamente",
			resultado:               50.0,
			primeiroDigito:          10.0,
			segundoDigito:           20.0,
			operador:                "+",
			input:                   "10.0\n+\n20.0\nsim\n+\n20\nnao\n",
			print: 					 casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Digitos inválidos devem ser identificados corretamente",
			resultado:               0.0,
			primeiroDigito:          0.0,
			segundoDigito:           0.0,
			operador:                "",
			input:                   "0.0\nfdgdfg\n0.0\n",
			print: 					 casosDeTestes,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			file, err := ioutil.TempFile("", "")
			if err != nil {
				teste.Fatal(err)
			}

			defer file.Close()

			_, err = io.WriteString(file, valorTeste.input)
			if err != nil {
				teste.Fatal(err)
			}

			_, err = file.Seek(0, os.SEEK_SET)
			if err != nil {
				teste.Fatal(err)
			}
			input := bufio.NewScanner(file)

			resultado, err := modoInterativo(valorTeste.primeiroDigito, valorTeste.segundoDigito, valorTeste.operador, input, valorTeste.print)

			if resultado != valorTeste.resultado {
				teste.Errorf("resultado recebido = %v , resultado esperado = %v", resultado, 50.0)
			}

			_, erro := modoInterativo(valorTeste.primeiroDigito, valorTeste.segundoDigito, valorTeste.operador, input, valorTeste.print)

			if erro == nil {
				teste.Errorf("erro recebido = %v , erro esperado = %v", erro, nil)
			}
		})
	}
}
func Test_modoExecução(teste *testing.T) {
	type parâmetrosRecebidos struct {
		numeros    []string
		operadores []string
		print      *os.File
	}

	casosDeTestes,err := os.Create("./testes/casos de teste da funcao modoExecução.txt")
	if err != nil{
		teste.Fatal(err)
	}

	testes := []struct {
		mensagemDeIdentificação string
		parâmetrosRecebidos     func(teste *testing.T) parâmetrosRecebidos
		print                   *os.File
		resultado               float64
	}{
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"10"},
					operadores: []string{"+"},
				}
			},
			resultado: 10.0,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Valor inteiro e negativo deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"-10"},
					operadores: []string{"+"},
				}
			},
			resultado: -10.0,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Float com . deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"10.1"},
					operadores: []string{"+"},
				}
			},
			resultado: 10.1,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Float com . e negativo deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"-10.1"},
					operadores: []string{"+"},
				}
			},
			resultado: -10.1,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Float com , deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"10,1"},
					operadores: []string{"+"},
				}
			},
			resultado: 10.1,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Digitos inválidos devem ser identificados corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"dsjfjshgfhgdfhfd"},
					operadores: []string{"+"},
				}
			},
			resultado: 0.0,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Operador inválido deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"0.0", "0,0"},
					operadores: []string{"fdgfgdfg"},
				}
			},
			resultado: 0.0,
			print: casosDeTestes,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			testeModoExecução := valorTeste.parâmetrosRecebidos(teste)

			valorRecebido := modoExecução(testeModoExecução.numeros, testeModoExecução.operadores, testeModoExecução.print)

			if !reflect.DeepEqual(valorRecebido, valorTeste.resultado) {
				teste.Errorf("modoExecução valorRecebido = %v,valorEsperado = %v", valorRecebido, valorTeste.resultado)
			}
		})
	}
}
func Test_lerInputs(teste *testing.T) {
	casosDeTestes,err := os.Create("./testes/casos de teste da funcao lerInputs.txt")
	if err != nil{
		teste.Fatal(err)
	}

	testes := []struct {
		mensagemDeIdentificação string
		primeiroDigito          string
		segundoDigito           string
		operador                string
		input                   string
		print                   *os.File
	}{
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			primeiroDigito:          "8",
			segundoDigito:           "2",
			operador:                "+",
			input:                   "8\n+\n2\nnao\n",
			print:                   casosDeTestes,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			file, err := ioutil.TempFile("", "")
			if err != nil {
				teste.Fatal(err)
			}

			defer file.Close()

			_, err = io.WriteString(file, valorTeste.input)
			if err != nil {
				teste.Fatal(err)
			}

			_, err = file.Seek(0, os.SEEK_SET)
			if err != nil {
				teste.Fatal(err)
			}
			input := bufio.NewScanner(file)

			primeiroDigito := lerInputs(input, "Digite o primeiro numero:", valorTeste.print)

			if primeiroDigito != valorTeste.primeiroDigito {
				teste.Errorf("digito recebido = %v ,digito esperado = %v ", primeiroDigito, valorTeste.primeiroDigito)
			}

			operador := lerInputs(input, "Digite o operador:", valorTeste.print)

			if operador != valorTeste.operador {
				teste.Errorf("digito recebido = %v ,digito esperado = %v ", operador, valorTeste.operador)
			}

			segundoDigito := lerInputs(input, "Digite o segundo numero:", valorTeste.print)

			if segundoDigito != valorTeste.segundoDigito {
				teste.Errorf("digito recebido = %v ,digito esperado = %v ", segundoDigito, valorTeste.segundoDigito)
			}

			novoCalculo := lerInputs(input, "Deseja fazer um novo calculo?", valorTeste.print)

			if novoCalculo != "nao" {
				teste.Errorf("digito recebido = %v ,digito esperado = %v ", novoCalculo, "nao")
			}
		})
	}
}
func Test_exibeErro(teste *testing.T) {
	type parâmetrosRecebidos struct {
		textoErro string
		print     *os.File
	}

	casosDeTestes,err := os.Create("./testes/casos de teste da funcao exibeErro.txt")
	if err != nil{
		teste.Fatal(err)
	}

	testes := []struct {
		mensagemDeIdentificação string
		parâmetrosRecebidos     func(teste *testing.T) parâmetrosRecebidos

		print                  *os.File
		mensagemDeErroEsperada string
	}{
		{
			mensagemDeIdentificação: "mensagem de erro deve ser identificada corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					textoErro: "Argumento inválido",
				}
			},
			mensagemDeErroEsperada: "Argumento inválido",
			print: casosDeTestes,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			testeExibeErro := valorTeste.parâmetrosRecebidos(teste)

			mensagemRecebida := exibeErro(testeExibeErro.textoErro, testeExibeErro.print)

			if !reflect.DeepEqual(mensagemRecebida, valorTeste.mensagemDeErroEsperada) {
				teste.Errorf("exibeErro mensagemRecebida = %v,mensagemEsperada = %v", mensagemRecebida, valorTeste.mensagemDeErroEsperada)
			}
		})
	}
}
func Test_tratarValor(teste *testing.T) {
	type parâmetrosRecebidos struct {
		valorDigitado string
		digito        string
		print         *os.File
	}

	casosDeTestes,err := os.Create("./testes/casos de teste da funcao tratarValor.txt")
	if err != nil{
		teste.Fatal(err)
	}

	testes := []struct {
		mensagemDeIdentificação string
		parâmetrosRecebidos     func(teste *testing.T) parâmetrosRecebidos

		print         *os.File
		valorEsperado float64
		erroEsperado  error
	}{
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					valorDigitado: "10",
					digito:        "primeiro digito",
				}
			},
			valorEsperado: 10.0,
			erroEsperado:  nil,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Float com , deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					valorDigitado: "10,1",
					digito:        "primeiro digito",
				}
			},
			valorEsperado: 10.1,
			erroEsperado:  nil,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Float com . deve ser identificado corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					valorDigitado: "10.1",
					digito:        "primeiro digito",
				}
			},
			valorEsperado: 10.1,
			erroEsperado:  nil,
			print: casosDeTestes,
		},
		{
			mensagemDeIdentificação: "Digitos inválidos devem ser identificados corretamente",
			parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					valorDigitado: "asdfasdfd2wqafdaq",
					digito:        "primeiro digito",
				}
			},
			valorEsperado: 0.0,
			print: casosDeTestes,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			testeTratarValor := valorTeste.parâmetrosRecebidos(teste)

			valorRecebido, err := tratarValor(testeTratarValor.valorDigitado, testeTratarValor.digito, testeTratarValor.print)

			if !reflect.DeepEqual(valorRecebido, valorTeste.valorEsperado) {
				teste.Errorf("tratarValor erro = %v, valorRecebido = %v, valorEsperado: %v", err, valorRecebido, valorTeste.valorEsperado)
			}
		})
	}
}
