package main // Define o pacote principal para um programa Go

import (
	"fmt" // Importa o pacote 'fmt' para formatação de saída
)

// Função para converter de Celsius para Kelvin
func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

// Função para converter de Kelvin para Celsius
func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

// Função para converter de Celsius para Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9/5) + 32
}

// Função para converter de Fahrenheit para Celsius
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5/9
}

// Função para converter de Kelvin para Fahrenheit
func kelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9/5 + 32
}

// Função para converter de Fahrenheit para Kelvin
func fahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)*5/9 + 273.15
}

func main() { // Função principal do programa
	var temperature float64 // Declaração de uma variável para armazenar a temperatura
	var choice int          // Declaração de uma variável para armazenar a escolha do usuário

	fmt.Println("Conversão de Temperaturas") // Exibe uma mensagem na saída padrão
	fmt.Println("1. Celsius para Kelvin")     // Exibe opções de conversão
	fmt.Println("2. Kelvin para Celsius")
	fmt.Println("3. Celsius para Fahrenheit")
	fmt.Println("4. Fahrenheit para Celsius")
	fmt.Println("5. Kelvin para Fahrenheit")
	fmt.Println("6. Fahrenheit para Kelvin")
	fmt.Print("Escolha a opção (1-6): ") // Solicita a escolha do usuário
	fmt.Scan(&choice)                   // Lê a escolha do usuário a partir da entrada padrão

	switch choice { // Inicia uma instrução switch baseada na escolha do usuário
	case 1: // Caso a escolha seja 1
		fmt.Print("Informe a temperatura em Celsius: ") // Solicita a temperatura em Celsius
		fmt.Scan(&temperature)                            // Lê a temperatura informada pelo usuário
		result := celsiusToKelvin(temperature)            // Chama a função de conversão
		fmt.Printf("%.2f Celsius é igual a %.2f Kelvin\n", temperature, result) // Exibe o resultado da conversão
	case 2: // Caso a escolha seja 2 (e assim por diante)
		fmt.Print("Informe a temperatura em Kelvin: ")
		fmt.Scan(&temperature)
		result := kelvinToCelsius(temperature)
		fmt.Printf("%.2f Kelvin é igual a %.2f Celsius\n", temperature, result)
	case 3:
		fmt.Print("Informe a temperatura em Celsius: ")
		fmt.Scan(&temperature)
		result := celsiusToFahrenheit(temperature)
		fmt.Printf("%.2f Celsius é igual a %.2f Fahrenheit\n", temperature, result)
	case 4:
		fmt.Print("Informe a temperatura em Fahrenheit: ")
		fmt.Scan(&temperature)
		result := fahrenheitToCelsius(temperature)
		fmt.Printf("%.2f Fahrenheit é igual a %.2f Celsius\n", temperature, result)
	case 5:
		fmt.Print("Informe a temperatura em Kelvin: ")
		fmt.Scan(&temperature)
		result := kelvinToFahrenheit(temperature)
		fmt.Printf("%.2f Kelvin é igual a %.2f Fahrenheit\n", temperature, result)
	case 6:
		fmt.Print("Informe a temperatura em Fahrenheit: ")
		fmt.Scan(&temperature)
		result := fahrenheitToKelvin(temperature)
		fmt.Printf("%.2f Fahrenheit é igual a %.2f Kelvin\n", temperature, result)
	default: // Caso a escolha não corresponda a nenhuma opção válida
		fmt.Println("Opção inválida. Por favor, escolha uma opção de 1 a 6.") // Exibe uma mensagem de erro
	}
}
