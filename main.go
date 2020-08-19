package main

import "fmt"

//PUNTEROS
func modificarVariable(param *string) {

	*param = "Cambio de valor"

}

func main() {
	//Punteros

	var curso = "Curso profesional de Go!"
	fmt.Println("Antes de la función: ", curso)

	modificarVariable(&curso) //referencia

	fmt.Println("Después de la función: ", curso)

}
