package main

import (
	"fmt"
	"os"
)

func main() {
	//RECOVER

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups!, al parecer el programa no finalizó correctamente")
		}
	}()

	if file, err := os.Open("hi.txt"); err != nil {

		panic("no fue posible obtener el archivo")

	} else {

		defer func() {
			fmt.Println("Cerramos el archivo")
			file.Close()
		}()

		contenido := make([]byte, 254)

		long, _ := file.Read(contenido)

		contenidoArchivo := string(contenido[0:long])

		fmt.Println(contenidoArchivo)
	}

}
