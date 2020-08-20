package globales

import "fmt"

var user string

func funcion1() {
	user = "Funcion1"
	fmt.Println("Funcion1: ", user)
}

func funcion2() {
	user = "Funcion2"
	fmt.Println("Funcion2: ", user)
}

func globales() {
	user = "Saúl"

	funcion1()
	funcion2()

	fmt.Println(user)
}
