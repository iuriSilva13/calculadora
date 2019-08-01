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
                                        primeiroValor: 0,
                                        segundoValor:  10,
                                        operador:      "+",
                                }
                        },
                        valorEsperado: 10.0,
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
