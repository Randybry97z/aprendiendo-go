package main

import "fmt"

func main() {
	//CONDICIONALES

	//if-else
	if false {
		fmt.Println("La condicion se cumple")
	} else {
		fmt.Println("La condicion no se cumple")
	}

	var edad int
	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad)

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
}
