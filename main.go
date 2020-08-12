package main

import "fmt"

func main() {
	//CONDICIONALES

	/*
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
	*/

	//else if
	var calificacion int
	fmt.Println("Ingresa tu calificación: ")
	fmt.Scanf("%d", &calificacion)

	if calificacion == 10 {
		fmt.Println("Tu calificación es perfecta")
	} else if calificacion == 8 || calificacion == 9 {
		fmt.Println("Tu calificación es excelente")
	} else if calificacion == 6 || calificacion == 7 {
		fmt.Println("Tu calificación es aprobatoria")
	} else if calificacion >= 0 && calificacion <= 5 {
		fmt.Println("Reprobaste")
	} else {
		fmt.Println("Calificación no válida")
	}
}
