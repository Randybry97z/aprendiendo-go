package main

func main() {
	/* //FOR

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	} */

	/* //WHILE (no existe pero se reemplaza con un for)

	numero := 12875
	contador := 0
	//while
	for numero > 0 {
		numero = numero / 10
		contador++
	}

	fmt.Println("La cantidad de digitos es: ", contador) */

	/* //FOREACH
	animales := [...]string{"Perro", "Gato", "Pez", "Vaca", "Cerdo"}

	// for index, element (si se quiere evitar el error de variable no usada, se coloca guión bajo)
	for _, element := range animales {
		fmt.Println("El valor es: ", element)
	} */

	/* //CONTINUE & BREAK

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue //Finaliza la iteración actual y pasa a la siguiente
		}

		if i == 8 {
			break //Finaliza la ejecución del ciclo actual
		}
		fmt.Println(i)
	} */
}
