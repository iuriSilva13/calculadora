package main

import (
        "reflect"
        "testing"
)

func Test_calcularValores(teste *testing.T) {
        type parâmetrosRecebidos struct {
				primeiroValor float64
				segundoValor  float64
				operador      string
        }
        testes := []struct {
                mensagemDeIdentificação string
                parâmetrosRecebidos func(teste *testing.T) parâmetrosRecebidos

                valorEsperado float64
                erroEsperado string
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
                        erroEsperado: "primeiro digito invalido",
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
                        erroEsperado: "primeiro digito invalido",
              },
        }

        for _, valorTeste := range testes {
                teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
                        testeCalcularValores := valorTeste.parâmetrosRecebidos(teste)

                        valorRecebido,err := calcularValores(testeCalcularValores.primeiroValor, testeCalcularValores.segundoValor,testeCalcularValores.operador)

                        if !reflect.DeepEqual(valorRecebido, valorTeste.valorEsperado) {
                                teste.Errorf("calcularValores erro = %v, valorRecebido = %v, valorEsperado: %v", err, valorRecebido, valorTeste.valorEsperado)
                        }
                })
        }
}
func Test_obterDadosDosInputs(teste *testing.T) {
        type parâmetrosRecebidos struct {
				primeiraVez bool
        }
        testes := []struct {
                mensagemDeIdentificação string
                parâmetrosRecebidos func(teste *testing.T) parâmetrosRecebidos

                primeiroValorEsperado float64
                segundoValorEsperado float64
                operador string
                erroEsperado error
        }{
                {
                        mensagemDeIdentificação: "True deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        primeiraVez: true,
                                }
                        },
                        primeiroValorEsperado: 0.0,
                        segundoValorEsperado: 0.0,
                        operador: "",
                        erroEsperado: nil,
                },
                {
                        mensagemDeIdentificação: "False deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        primeiraVez: false,
                                }
                        },
                        primeiroValorEsperado: 0.0,
                        segundoValorEsperado: 0.0,
                        operador: "",
                        erroEsperado: nil,
                },
        }

        for _, valorTeste := range testes {
                teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
                        testeCalcularValores := valorTeste.parâmetrosRecebidos(teste)

                        primeiroValor,segundoValor,operator,err := obterDadosDosInputs(testeCalcularValores.primeiraVez)

                        if !reflect.DeepEqual(primeiroValor,valorTeste.primeiroValorEsperado) {
        teste.Errorf("obterDadosDosInputs primeiroValorRecebido = %v,segundoValorRecebido = %v,operadorRecebido = %v,erroRecebido = %v,primeiroValorEsperado = %v,segundoValorEsperado = %v,operadorEsperado = %v,erroEsperado = %v", primeiroValor,segundoValor,operator,err,valorTeste.primeiroValorEsperado,valorTeste.segundoValorEsperado,valorTeste.operador,valorTeste.erroEsperado)
                        }
                })
        }
}
func Test_modoInterativo(teste *testing.T) {
        type parâmetrosRecebidos struct {
                primeiroDigito float64
                segundoDigito float64
                operador string
        }
        testes := []struct {
                mensagemDeIdentificação string
                parâmetrosRecebidos func(teste *testing.T) parâmetrosRecebidos

                valorEsperado float64
                erroEsperado error
        }{
                {
                        mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        primeiroDigito: 0,
                                        segundoDigito: 0,
                                        operador: "",
                                }
                        },
                       valorEsperado: 0.0,
                       erroEsperado: nil,
                },
                {
                        mensagemDeIdentificação: "Float deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        primeiroDigito: 0.0,
                                        segundoDigito: 0.0,
                                        operador: "",
                                }
                        },
                       valorEsperado: 0.0,
                       erroEsperado: nil,
                },
        }

        for _, valorTeste := range testes {
                teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
                        testeCalcularValores := valorTeste.parâmetrosRecebidos(teste)

                        valorRecebido,err := modoInterativo(testeCalcularValores.primeiroDigito,testeCalcularValores.segundoDigito,testeCalcularValores.operador)

                        if !reflect.DeepEqual(valorRecebido,valorTeste.valorEsperado) {
                                teste.Errorf("modoInterativo erro = %v,valorRecebido = %v,valorEsperado = %v",err,valorRecebido,valorTeste.valorEsperado)
                        }
                })
        }
}
func Test_modoExecução(teste *testing.T) {
        type parâmetrosRecebidos struct {
                numeros []string
                operadores []string
        }
        testes := []struct {
                mensagemDeIdentificação string
                parâmetrosRecebidos func(teste *testing.T) parâmetrosRecebidos

                resultado float64
        }{
                {
                        mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        numeros: []string{"10"},
                                        operadores: []string{"+"},
                                }
                        },
                       resultado: 10.0,
                },
                {
                        mensagemDeIdentificação: "Valor inteiro e negativo deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        numeros: []string{"-10"},
                                        operadores: []string{"+"},
                                }
                        },
                       resultado: -10.0,
                },
                {
                        mensagemDeIdentificação: "Float com . deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        numeros: []string{"10.1"},
                                        operadores: []string{"+"},
                                }
                        },
                       resultado: 10.1,
                },
                {
                        mensagemDeIdentificação: "Float com . e negativo deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        numeros: []string{"-10.1"},
                                        operadores: []string{"+"},
                                }
                        },
                       resultado: -10.1,
                },
                {
                        mensagemDeIdentificação: "Float com , deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        numeros: []string{"10,1"},
                                        operadores: []string{"+"},
                                }
                        },
                       resultado: 10.1,
                },
                {
                        mensagemDeIdentificação: "Digitos inválidos devem ser identificados corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        numeros: []string{"dsjfjsdfhfd"},
                                        operadores: []string{"+"},
                                }
                        },
                       resultado: 0.0,
                },
        }

        for _, valorTeste := range testes {
                teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
                        testeCalcularValores := valorTeste.parâmetrosRecebidos(teste)

                        valorRecebido := modoExecução(testeCalcularValores.numeros,testeCalcularValores.operadores)

                        if !reflect.DeepEqual(valorRecebido,valorTeste.resultado) {
                                teste.Errorf("modoExecução valorRecebido = %v,valorEsperado = %v",valorRecebido,valorTeste.resultado)
                        }
                })
        }
}
func Test_exibeErro(teste *testing.T) {
        type parâmetrosRecebidos struct {
                textoErro string
        }
        testes := []struct {
                mensagemDeIdentificação string
                parâmetrosRecebidos func(teste *testing.T) parâmetrosRecebidos

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
                },
        }

        for _, valorTeste := range testes {
                teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
                        testeMensagmErro := valorTeste.parâmetrosRecebidos(teste)

                        mensagemRecebida := exibeErro(testeMensagmErro.textoErro)

                        if !reflect.DeepEqual(mensagemRecebida,valorTeste.mensagemDeErroEsperada) {
                                teste.Errorf("exibeErro mensagemRecebida = %v,mensagemEsperada = %v",mensagemRecebida,valorTeste.mensagemDeErroEsperada)
                        }
                })
        }
}
func Test_tratarValor(teste *testing.T) {
        type parâmetrosRecebidos struct {
                valorDigitado string
                digito        string
        }
        testes := []struct {
                mensagemDeIdentificação string
                parâmetrosRecebidos func(teste *testing.T) parâmetrosRecebidos

                valorEsperado float64
                erroEsperado error
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
                        erroEsperado: nil,
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
                        erroEsperado: nil,
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
                        erroEsperado: nil,
                },
                {
                        mensagemDeIdentificação: "Valores devem ser identificados corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        valorDigitado: "asdfasdfd2wqafdaq",
                                        digito:        "primeiro digito",
                                }
                        },
                        valorEsperado: 0.0,
                        erroEsperado:nil,
                },
        }

        for _, valorTeste := range testes {
                teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
                        testeTratarValor := valorTeste.parâmetrosRecebidos(teste)

                        valorRecebido,err := tratarValor(testeTratarValor.valorDigitado, testeTratarValor.digito)

                        if !reflect.DeepEqual(valorRecebido, valorTeste.valorEsperado) {
                                teste.Errorf("tratarValor erro = %v, valorRecebido = %v, valorEsperado: %v", err, valorRecebido, valorTeste.valorEsperado)
                        }
                })
        }
}
