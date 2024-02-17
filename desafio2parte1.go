// Desafio 2 parte1: criando um programa que trabalhe com o operador % e for:
// um código que exiba todos os números entre 1 e 100 que são divisíveis por 3.

package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, "é divisível por 3")
		}
	}
}
