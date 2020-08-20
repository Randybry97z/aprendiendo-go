package funVariables

import "fmt"

//FUNCIONES COMO VARIABLES
func factorial(num int) int {

	if num == 1 {
		return 1
	}

	return factorial(num-1) * num
}

type customFunction func(n int) int

func funVariables() {

	//Funciones como variables
	//var miFuncion = factorial
	var miFuncion customFunction

	if miFuncion == nil {
		miFuncion = factorial
	}

	result := miFuncion(3)

	fmt.Println(result)
}
