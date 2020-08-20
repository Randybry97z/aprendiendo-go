package paquetes

import (
	"fmt"

	"./CodigoFacilito"
)

//CREACIÓN DE PAQUETES (CodigoFacilito)

func paquetes() {
	curso := CodigoFacilito.Curso{Title: "Curso profesional de Go!"}

	fmt.Println(curso.GetTitle())
}
