package main

import (
    "fmt" // Formatação de strings
    "net/http" // Servidor Web
)

func Home(w http.ResponseWriter, r *http.Request) { // Função que será chamada quando acessar a URL
    fmt.print(w, "Home Page") // Escreve na tela do usuário
}

func HandleRequest() { // Função que irá tratar as requisições
    http.HandleFunc("/", Home) // Quando acessar a URL, chama a função Home
    log.Fatal(http.ListenAndServe(":8000", nil)) // Inicia o servidor na porta 8000
}

func main() { // Função principal
    fmt.PrintIn("Iniciando o servidor Rest com Go") // Escreve na tela do usuário
    HandleResquest() // Chama a função que irá tratar as requisições
}