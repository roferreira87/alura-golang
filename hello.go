package main

import "fmt"

func main() {
	nome := "Rogério"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("ESte programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("O comando escolhido foi", comando)
}
