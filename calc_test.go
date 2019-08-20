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
                                        primeiroValor: 10,
                                        segundoValor:  5,
                                        operador:      "-",
                                }
                        },
                        valorEsperado: 5.0,
                },
                {
                        mensagemDeIdentificação: "Float com . deve ser identificado corretamente",
                        parâmetrosRecebidos: func(*testing.T) parâmetrosRecebidos {
                              return parâmetrosRecebidos{
                                      primeiroValor: 10.5,
                                      segundoValor:  7.1,
                                      operador:      "-",
                              }
                        },
                        valorEsperado: 3.4000000000000004,
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
