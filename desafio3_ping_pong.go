'''
Aqui está um exemplo simples de um programa Go que: 
usa goroutines e canais para exibir "ping" e "pong" alternadamente. 

Este programa cria três goroutines. 
As funções `ping` e `pong` enviam as strings "ping" e "pong" para o canal `c` indefinidamente. 
A função `printer` recebe uma string do canal `c`, imprime-a e, em seguida, dorme por um segundo. 
Este processo é repetido indefinidamente. 
A função `main` inicia as goroutines e, em seguida, aguarda a entrada do usuário para terminar o programa.
'''

package main

import (
	"fmt" // Importa o pacote fmt para imprimir mensagens na tela
	"time" // Importa o pacote time para usar funções relacionadas ao tempo
)

func ping(c chan string) { // Função ping que recebe um canal de string como argumento
	for i := 0; ; i++ { // Loop infinito que incrementa a variável i a cada iteração 
		c <- "ping" // Envia a string "ping" para o canal c 
	}
}

func pong(c chan string) { // Função pong que recebe um canal de string como argumento 
	for i := 0; ; i++ { // Loop infinito que incrementa a variável i a cada iteração
		c <- "pong" // Envia a string "pong" para o canal c
	}
}

func printer(c chan string) { // Função printer que recebe um canal de string como argumento
	for {
		msg := <-c // Recebe uma string do canal c e armazena na variável msg 
		fmt.Println(msg) // Imprime a string msg na tela 
		time.Sleep(time.Second * 1) // Dorme por um segundo 
	}
}

func main() { // Função principal do programa 
	var c chan string = make(chan string) // Cria um canal de string e armazena na variável c 

	go ping(c) // Inicia a goroutine ping, passando o canal c como argumento
	go pong(c) // Inicia a goroutine pong, passando o canal c como argumento
	go printer(c) // Inicia a goroutine printer, passando o canal c como argumento

	var input string // Declara uma variável input do tipo string
	fmt.Scanln(&input) // Aguarda a entrada do usuário para encerrar o programa
}
