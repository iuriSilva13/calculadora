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
                },
                {
                        mensagemDeIdentificação: "Float com . deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                              return parâmetrosRecebidos{
                                      primeiroValor: 10.520,
                                      segundoValor:  5.200,
                                      operador:      "*",
                              }
                        },
                        valorEsperado: 54.704,
              },
        }

        for _, valorTeste := range testes {
                teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
                        testeCalcularValores := valorTeste.parâmetrosRecebidos(teste)

                        valorRecebido := calcularValores(testeCalcularValores.primeiroValor, testeCalcularValores.segundoValor,testeCalcularValores.operador)

                        if !reflect.DeepEqual(valorRecebido, valorTeste.valorEsperado) {
                                teste.Errorf("calcularValores valorRecebido = %v, valorEsperado: %v", valorRecebido, valorTeste.valorEsperado)
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
        }{
                {
                        mensagemDeIdentificação: "Inteiro deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                                return parâmetrosRecebidos{
                                        valorDigitado: "10",
                                        digito:        "primeiro",
                                }
                        },
                        valorEsperado: 10.0,
                },
        }

        for _, valorTeste := range testes {
                teste.Run(valorTeste.mensagemDeIdentificação, func(teste *testing.T) {
                        testeTratarValor := valorTeste.parâmetrosRecebidos(teste)

                        valorRecebido := tratarValor(testeTratarValor.valorDigitado, testeTratarValor.digito)

                        if !reflect.DeepEqual(valorRecebido, valorTeste.valorEsperado) {
                                teste.Errorf("tratarValor valorRecebido = %v, valorEsperado: %v", valorRecebido, valorTeste.valorEsperado)
                        }
                })
        }
}
