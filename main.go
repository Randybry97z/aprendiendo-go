package main

import "fmt"

/* //RETORNAR FUNCIONES

type Operacion func(balance, cantidad int) int

func crearOperacion(tipo string) Operacion {

	if tipo == "suma" {
		return func(balance, cantidad int) int { return balance + cantidad }
	} else if tipo == "resta" {
		return func(balance, cantidad int) int { return balance - cantidad }
	} else {
		return func(balance, cantidad int) int { return balance * cantidad }
	}

} */

func main() { //Bloque 1

	/* //Retornar funciones
	suma := crearOperacion("suma")

	resultado := suma(40, 50)

	fmt.Println("El resultado es: ", resultado) */

	//BLOQUES Y ALCANCES
	//Una variable sólo vive en su bloque
	var curso = "Curso de Go!"
	var version = 10

	fmt.Println(curso)
	fmt.Println(version, "\n")

	{ //Bloque 2
		var version = 5

		fmt.Println("Bloque 2")
		fmt.Println(curso)
		fmt.Println(version)

	}

}
