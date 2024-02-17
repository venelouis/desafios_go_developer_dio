//  criar um código que sempre que aparecer um número multiplo de 3 ele escreva "Pin"
//  sempre que aparecer um número multiplo de 5 ele escreva "Pan"
// O programa deve imprimir numeros de 1 a 100, mas escrever "Pin" e "Pan" para os multiplos de 3 e 5 respectivamente.
// e "PinPan" para os multiplos de 3 e 5 ao mesmo tempo.

package main

func main() {

	for i := 1; i <= 100; i++ { // para i de 1 até 100 faça i++ (i = i + 1) (i += 1)
		if i%3 == 0 { // se o resto da divisão de i por 3 for igual a 0
			if i%5 == 0 { // se o resto da divisão de i por 5 for igual a 0
				println("PinPan") // imprima "PinPan"
			} else { // senão
				println("Pin") // imprima "Pin"
			} // fim do if
		} else if i%5 == 0 { // senão se o resto da divisão de i por 5 for igual a 0
			println("Pan") // imprima "Pan"
		} else { // senão
			println(i) // imprima i
		} // fim do if
	} // fim do for
} // fim do main
