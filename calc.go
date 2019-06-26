package main

import (
	"fmt"
)

func main() {
	var primeiroNumero , segundoNumero float64
	var operação string

	fmt.Print("digite um numero:")
	fmt.Scan(&primeiroNumero)

	fmt.Print("digite a operação:")
	fmt.Scan(&operação)

	fmt.Print("digite outro numero:")
	fmt.Scan(&segundoNumero)
}