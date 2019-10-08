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
	os.Mkdir("output",0777)
	output := m.Run()
	os.RemoveAll("output")
	os.Exit(output)
}
func Test_calcularValores(teste *testing.T) {
	type parâmetrosRecebidos struct {
		primeiroValor float64
		segundoValor  float64
		operador      string
	}

	file,err := os.Create("./output/calcularValores.txt")
	if err != nil{
		teste.Fatal(err)
	}

	defer file.Close()

	testes := []struct {
		mensagemDeIdentificação string
		dadosRecebidos     func(teste *testing.T) parâmetrosRecebidos

		w         *os.File
		valorEsperado float64
		erroEsperado  string
	}{
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroValor: -10,
					segundoValor:  -5,
					operador:      "+",
				}
			},
			valorEsperado: -15.0,
			w: file,
		},
		{
			mensagemDeIdentificação: "Float com . deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroValor: -80.3,
					segundoValor:  -4.2,
					operador:      "/",
				}
			},
			valorEsperado: 19.119047619047617,
			w: file,
		},
		{
			mensagemDeIdentificação: "Operador de subtração deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroValor: 45,
					segundoValor:  35,
					operador:      "-",
				}
			},
			valorEsperado: 10.0,
			w: file,
		},
		{
			mensagemDeIdentificação: "Operador de multiplicação deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroValor: 10,
					segundoValor:  5,
					operador:      "*",
				}
			},
			valorEsperado: 50.0,
			w: file,
		},
		{
			mensagemDeIdentificação: "Operador deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroValor: 10,
					segundoValor:  45,
					operador:      "gfdgfdgfd",
				}
			},
			valorEsperado: 0.0,
			erroEsperado:  "Argumento inválido",
			w: file,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			testeCalcularValores := valorTeste.dadosRecebidos(teste)

			valorRecebido, err := calcularValores(testeCalcularValores.primeiroValor, testeCalcularValores.segundoValor, testeCalcularValores.operador, valorTeste.w)

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
		w          *os.File
	}

	file,err := os.Create("./output/validarEntradas.txt")
	if err != nil{
		teste.Fatal(err)
	}

	defer file.Close()

	testes := []struct {
		mensagemDeIdentificação string
		dadosRecebidos     func(teste *testing.T) parâmetrosRecebidos

		w                 *os.File
		primeiroValorEsperado float64
		segundoValorEsperado  float64
		primeiraVez           bool
		erroEsperado          error
	}{
		{
			mensagemDeIdentificação: "caso de erro com valor false deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroDigito: "",
					segundoDigito:  "",
					primeiraVez:    false,
				}
			},
			primeiroValorEsperado: 0.0,
			segundoValorEsperado:  0.0,
			primeiraVez:		   false,
			w: file,
		},
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
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
			w: file,
		},
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
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
			w: file,
		},
		{
			mensagemDeIdentificação: "caso de erro com valor true deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					primeiroDigito: "",
					segundoDigito:  "",
					primeiraVez:    true,
				}
			},
			primeiroValorEsperado: 0.0,
			segundoValorEsperado:  0.0,
			primeiraVez:           true,
			w: file,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			testeValidarEntradas := valorTeste.dadosRecebidos(teste)

			primeiroValor, segundoValor, err := validarEntradas(testeValidarEntradas.primeiroDigito, testeValidarEntradas.segundoDigito, testeValidarEntradas.primeiraVez, testeValidarEntradas.w)

			if !reflect.DeepEqual(segundoValor, valorTeste.segundoValorEsperado) {
				teste.Errorf("validarEntradas primeiroValorRecebido = %v,segundoValorRecebido = %v,primeiroCalculoRecebido = %v,primeiroValorEsperado = %v,segundoValorEsperado = %v, erroEsperado = %v", primeiroValor, segundoValor, err, valorTeste.primeiroValorEsperado, valorTeste.segundoValorEsperado, valorTeste.erroEsperado)
			}
		})
	}
}
func Test_obterDadosDosInputs(teste *testing.T) {
	file,err := os.Create("./output/obterDadosDosInputs.txt")
	if err != nil{
		teste.Fatal(err)
	}

	defer file.Close()

	testes := []struct {
		mensagemDeIdentificação string
		primeiraVez             bool
		primeiroDigito          float64
		segundoDigito           float64
		operador                string
		err                     error
		input                   string
		w                   *os.File
	}{
		{
			mensagemDeIdentificação: "Float64 deve ser identificado corretamente",
			primeiraVez:             false,
			primeiroDigito:          25.0,
			segundoDigito:           4.0,
			operador:                "+",
			err:                     nil,
			input:                   "25.0\n+\n4.0\n",
			w: file,
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

			primeiroDigito, segundoDigito, operador, err := obterDadosDosInputs(valorTeste.primeiraVez, input, valorTeste.w)

			if valorTeste.primeiraVez != false {
				teste.Errorf("primeiro digito recebido = %v ,segundo digito recebido = %v,operador recebido = %v, primeiro digito esperado = %v, segundo digito esperado = %v,operador esperado = %v ", primeiroDigito, segundoDigito, operador, valorTeste.primeiroDigito, valorTeste.segundoDigito, valorTeste.operador)
			}
		})
	}
}
func Test_modoInterativo(teste *testing.T) {
	file,err := os.Create("./output/modoInterativo.txt")
	if err != nil{
		teste.Fatal(err)
	}

	defer file.Close()

	testes := []struct {
		mensagemDeIdentificação string
		resultado               float64
		primeiroDigito          float64
		segundoDigito           float64
		operador                string
		input                   string
		w                   *os.File
	}{
		{
			mensagemDeIdentificação: "Float64 deve ser identificado corretamente",
			resultado:               50.0,
			primeiroDigito:          10.0,
			segundoDigito:           20.0,
			operador:                "+",
			input:                   "10.0\n+\n20.0\nsim\n+\n20\nnao\n",
			w: 					 file,
		},
		{
			mensagemDeIdentificação: "Digitos inválidos devem ser identificados corretamente",
			resultado:               0.0,
			primeiroDigito:          0.0,
			segundoDigito:           0.0,
			operador:                "",
			input:                   "0.0\nfdgdfg\n0.0\n",
			w: 					 file,
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

			resultado, err := modoInterativo(valorTeste.primeiroDigito, valorTeste.segundoDigito, valorTeste.operador, input, valorTeste.w)

			if resultado != valorTeste.resultado {
				teste.Errorf("resultado recebido = %v , resultado esperado = %v", resultado, 50.0)
			}

			_, erro := modoInterativo(valorTeste.primeiroDigito, valorTeste.segundoDigito, valorTeste.operador, input, valorTeste.w)

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
		w      *os.File
	}

	file,err := os.Create("./output/modoExecução.txt")
	if err != nil{
		teste.Fatal(err)
	}

	defer file.Close()

	testes := []struct {
		mensagemDeIdentificação string
		dadosRecebidos     func(teste *testing.T) parâmetrosRecebidos
		w                   *os.File
		resultado               float64
	}{
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"10"},
					operadores: []string{"+"},
				}
			},
			resultado: 10.0,
			w: file,
		},
		{
			mensagemDeIdentificação: "Valor inteiro e negativo deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"-10"},
					operadores: []string{"+"},
				}
			},
			resultado: -10.0,
			w: file,
		},
		{
			mensagemDeIdentificação: "Float com . deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"10.1"},
					operadores: []string{"+"},
				}
			},
			resultado: 10.1,
			w: file,
		},
		{
			mensagemDeIdentificação: "Float com . e negativo deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"-10.1"},
					operadores: []string{"+"},
				}
			},
			resultado: -10.1,
			w: file,
		},
		{
			mensagemDeIdentificação: "Float com , deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"10,1"},
					operadores: []string{"+"},
				}
			},
			resultado: 10.1,
			w: file,
		},
		{
			mensagemDeIdentificação: "Digitos inválidos devem ser identificados corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"dsjfjshgfhgdfhfd"},
					operadores: []string{"+"},
				}
			},
			resultado: 0.0,
			w: file,
		},
		{
			mensagemDeIdentificação: "Operador inválido deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					numeros:    []string{"0.0", "0,0"},
					operadores: []string{"fdgfgdfg"},
				}
			},
			resultado: 0.0,
			w: file,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			testeModoExecução := valorTeste.dadosRecebidos(teste)

			valorRecebido := modoExecução(testeModoExecução.numeros, testeModoExecução.operadores, testeModoExecução.w)

			if !reflect.DeepEqual(valorRecebido, valorTeste.resultado) {
				teste.Errorf("modoExecução valorRecebido = %v,valorEsperado = %v", valorRecebido, valorTeste.resultado)
			}
		})
	}
}
func Test_lerInputs(teste *testing.T) {
	file,err := os.Create("./output/lerInputs.txt")
	if err != nil{
		teste.Fatal(err)
	}

	defer file.Close()

	testes := []struct {
		mensagemDeIdentificação string
		primeiroDigito          string
		segundoDigito           string
		operador                string
		input                   string
		w                   *os.File
	}{
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			primeiroDigito:          "8",
			segundoDigito:           "2",
			operador:                "+",
			input:                   "8\n+\n2\nnao\n",
			w:                   file,
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

			primeiroDigito := lerInputs(input, "Digite o primeiro numero:", valorTeste.w)

			if primeiroDigito != valorTeste.primeiroDigito {
				teste.Errorf("digito recebido = %v ,digito esperado = %v ", primeiroDigito, valorTeste.primeiroDigito)
			}

			operador := lerInputs(input, "Digite o operador:", valorTeste.w)

			if operador != valorTeste.operador {
				teste.Errorf("digito recebido = %v ,digito esperado = %v ", operador, valorTeste.operador)
			}

			segundoDigito := lerInputs(input, "Digite o segundo numero:", valorTeste.w)

			if segundoDigito != valorTeste.segundoDigito {
				teste.Errorf("digito recebido = %v ,digito esperado = %v ", segundoDigito, valorTeste.segundoDigito)
			}

			novoCalculo := lerInputs(input, "Deseja fazer um novo calculo?", valorTeste.w)

			if novoCalculo != "nao" {
				teste.Errorf("digito recebido = %v ,digito esperado = %v ", novoCalculo, "nao")
			}
		})
	}
}
func Test_exibeErro(teste *testing.T) {
	type parâmetrosRecebidos struct {
		textoErro string
		w     *os.File
	}

	file,err := os.Create("./output/exibeErro.txt")
	if err != nil{
		teste.Fatal(err)
	}

	defer file.Close()

	testes := []struct {
		mensagemDeIdentificação string
		dadosRecebidos     func(teste *testing.T) parâmetrosRecebidos

		w                  *os.File
		mensagemDeErroEsperada string
	}{
		{
			mensagemDeIdentificação: "mensagem de erro deve ser identificada corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					textoErro: "Argumento inválido",
				}
			},
			mensagemDeErroEsperada: "Argumento inválido",
			w: file,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			testeExibeErro := valorTeste.dadosRecebidos(teste)

			mensagemRecebida := exibeErro(testeExibeErro.textoErro, testeExibeErro.w)

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
		w         *os.File
	}

	file,err := os.Create("./output/tratarValor.txt")
	if err != nil{
		teste.Fatal(err)
	}

	defer file.Close()

	testes := []struct {
		mensagemDeIdentificação string
		dadosRecebidos     func(teste *testing.T) parâmetrosRecebidos

		w         *os.File
		valorEsperado float64
		erroEsperado  error
	}{
		{
			mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					valorDigitado: "10",
					digito:        "primeiro digito",
				}
			},
			valorEsperado: 10.0,
			erroEsperado:  nil,
			w: file,
		},
		{
			mensagemDeIdentificação: "Float com , deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					valorDigitado: "10,1",
					digito:        "primeiro digito",
				}
			},
			valorEsperado: 10.1,
			erroEsperado:  nil,
			w: file,
		},
		{
			mensagemDeIdentificação: "Float com . deve ser identificado corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					valorDigitado: "10.1",
					digito:        "primeiro digito",
				}
			},
			valorEsperado: 10.1,
			erroEsperado:  nil,
			w: file,
		},
		{
			mensagemDeIdentificação: "Digitos inválidos devem ser identificados corretamente",
			dadosRecebidos: func(*testing.T) parâmetrosRecebidos {
				return parâmetrosRecebidos{
					valorDigitado: "asdfasdfd2wqafdaq",
					digito:        "primeiro digito",
				}
			},
			valorEsperado: 0.0,
			w: file,
		},
	}

	for _, valorTeste := range testes {
		teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
			testeTratarValor := valorTeste.dadosRecebidos(teste)

			valorRecebido, err := tratarValor(testeTratarValor.valorDigitado, testeTratarValor.digito, testeTratarValor.w)

			if !reflect.DeepEqual(valorRecebido, valorTeste.valorEsperado) {
				teste.Errorf("tratarValor erro = %v, valorRecebido = %v, valorEsperado: %v", err, valorRecebido, valorTeste.valorEsperado)
			}
		})
	}
}
