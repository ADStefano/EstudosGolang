package main

import "fmt"


func main()  {
	
	// Dois jeitos simples de se fazer um for

	// Com apenas uma condição
	i := 1
	for i <= 4{
		fmt.Println("Valor de i:",i)
		i ++
	}

	// Inicializando a variável/Condição/Após o for
	for j := 2; j < 6; j++{
		fmt.Println("Valor de j:",j)
	}
}