package main

import (
	"fmt"
	"os"
)

//PROGRAMAR FUNCIONES
/* func funcion1() {
	fmt.Println("Hola soy la primera función")
}

func funcion2() {
	fmt.Println("Hola soy la segunda función")
} */

func main() {
	//Programar funciones

	/* //defer programa una funcion para que se ejecute cuando el bloque finalice
	defer funcion1()
	funcion2() */
	file, err := os.Open("example.txt")

	if err != nil {
		panic("no fue posible obtener el archivo")
	}

	defer func() {
		fmt.Println("Cerramos el archivo")
		file.Close()
	}()

	contenido := make([]byte, 254)

	long, _ := file.Read(contenido)

	contenidoArchivo := string(contenido[0:long])

	fmt.Println(contenidoArchivo)
}
