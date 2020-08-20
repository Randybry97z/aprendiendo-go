package main

import (
	"fmt"

	"./CodigoFacilito"
)

//CREACIÓN DE PAQUETES (CodigoFacilito)

func main() {
	curso := CodigoFacilito.Curso{Title: "Curso profesional de Go!"}

	fmt.Println(curso.GetTitle())
}
