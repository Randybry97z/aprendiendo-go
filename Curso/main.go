package main

import (
	"fmt"

	"./CodigoFacilito"
)

//MODIFICADORES DE ACCESO (si el método inicia con mayúsculas será público, de lo contrario será privado)

func main() {
	curso := CodigoFacilito.Curso{Title: "Curso profesional de Go!"}

	fmt.Println(curso.ObtenerTitulo())
}
