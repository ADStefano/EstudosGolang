package main // O package main avisa o compilador que ele deve ser compilado como executável e não como uma biblioteca compartilhada 

import ( 
	"fmt" 
	"teste/calculadora" // teste é o nome do módulo e calculadora o nome da pasta
)

func main()  { 
	
	// Quando declarada a variável DEVE ser utilizada
	//Métodos de se declarar uma variável:

	// A sintaxe := declara e inicializa a variável
	a := 4 

	// var declara uma ou mais variáveis
	var b, c int = 2, 3 
	fmt.Printf("Valor de b: %v e valor de c: %v\n",b,c)

	// Quando não se atribui um valor à variável ela recebe o valor 0
	var x int 
	fmt.Println("Int sem valor:",x)

	// No caso das Strings ela recebe um valor em branco ""
	var y string
	fmt.Println("String sem valor:",y)
	
	// Já boolean recebe como valor false
	var z bool 
	fmt.Println("Boolean sem valor:",z) // Utilizando o fmt.Printf para formatar a string

	// const é utilizado para declarar uma constante
	const con string = "Constante"
	fmt.Println(con)

	// Chamando as funções do importadas do teste/calculadora, para uma função poder ser importada o nome precisa começar com letra maiúscula
	divisao := calculadora.Dividir(a,b)
	soma := calculadora.Somar(a,b)
	multiplicacao := calculadora.Multiplicar(a,b)

	fmt.Println(divisao)
	fmt.Println(soma)
	fmt.Println(multiplicacao)

	
	
}