package errores

import (
	"errors"
	"fmt"
)

//Manejo de errores
func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir sobre cero")
	} else {
		return dividendo / divisor, nil
	}
}

func errores() {

	//MANEJO DE ERRORES
	if resultado, err := division(10, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El resultado es: ", resultado)
	}
}
