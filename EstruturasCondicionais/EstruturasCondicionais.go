package main

import "fmt"

func main()  {

	num := 4
	num1 := 8

	//No Go não há operador ternário então é necessário utilizar o if else até para as condições mais básicas

	//É possível ter um if sem else
	if num != num1{
		fmt.Printf("%v é diferente de %v",num,num1)
	}
	
	// Cláusula if else normal
	if num1%4 == 0{
		fmt.Printf("%v é divisível por 4",num1)
	} else{
		fmt.Printf("%v não é divisível por 4",num1)
	}

	fmt.Print("\n")

	// Utiliza-se o else if se o código precisar de mais condições antes do else
	if num < num1{
		fmt.Printf("%v é menor que %v",num,num1)
	} else if num == num1{
		fmt.Printf("%v é igual a %v",num,num1)
	} else{
		fmt.Printf("%v é maior que %v",num,num1)
	}


}