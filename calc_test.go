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
