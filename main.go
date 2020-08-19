package main

import "fmt"

/*
//DEFINICIÓN DE FUNCIONES

func saludo(user string) {
	fmt.Println("Hola ", user)
}

func add(numero1, numero2 int) int {
	return numero1 + numero2
}

//Retornar diferentes tipos en funcion, se definen entre paréntesis
func add(numero1, numero2 int) (int, string) {
	return numero1 + numero2, "Un mensaje desde suma"
} */

//FUNCIONES POR ARGUMENTO

//Creación de tipo de dato
type Operacion func(balance, cantidad int) int

func suma(num1, num2 int) int { //Función de tipo Operacion
	return num1 + num2
}

func resta(num1, num2 int) int { //Función de tipo Operacion
	return num1 - num2
}

func ejecutarOperacion(funcion Operacion, balance, cantidad int) {
	fmt.Println("Antes de la operación")

	result := funcion(balance, cantidad)
	fmt.Println("El resultado es: ", result)

	fmt.Println("Despues de la operación")
}

func main() {

	//	result, message := add(2, 5)
	//	fmt.Println(result, message)

	/* 	//FUNCIÓN ANÓNIMA
	func() {
		fmt.Println("Hola baby")
	}()

	myFunc := func(user string) string {
		message := fmt.Sprintf("Hola %s, te saludamos desde función sin nombre", user)
		return message
	}

	finalMessage := myFunc("Saúl")
		 fmt.Println(finalMessage) */
	ejecutarOperacion(suma, 10, 20)

}
