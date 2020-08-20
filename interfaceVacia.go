package ivacias

import "fmt"

//INTERFACES VACIAS
func mostrarVariable(objeto interface{}) {
	fmt.Printf("El valor de la variable es: %v\n", objeto)
}

func suma(num1, num2 int) int {
	return num1 + num2
}

type User struct {
	Nombre string
}

func ivacias() {

	usuario := User{Nombre: "Eduardo"}
	mostrarVariable(usuario)

}
