package main

func main() {
	/*
		//ARRAYS

		var nums [5]int

		nums[0] = 100
		nums[1] = 200
		nums[2] = 300
		nums[3] = 400
		nums[4] = 500 //Sólo se puede asignar un mismo tipo de dato a los elementos
		//nums := [5]int{100, 200, 300, 400, 500}

		nums := [...]int{100, 200, 300, 400, 500}

		fmt.Println(nums[2])

		monedas := [...]string{0: "Peso mexicano", 1: "Dólar", 2: "Euro"}

		//sub array (slice)
		subMonedas := monedas[0:3]
		fmt.Println(subMonedas)
	*/

	/*
			//SLICE (se puede aumentar o disminuir su longitud)

			numeros := []int{1, 2, 3, 4}

			numeros = append(numeros, 5)
			numeros = append(numeros, 6)
			numeros = append(numeros, 7)
			numeros = append(numeros, 8)
		numeros = append(numeros, 9)
		numeros = append(numeros, 10)

		nuevoSlice := numeros[0:5]
		tercerSlice := numeros[0:3]

		numeros[0] = 100

		fmt.Println(numeros)
		fmt.Println(nuevoSlice)
		fmt.Println(tercerSlice)

		meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre"}
		//Se debe tener en cuenta: Un puntero Una longitud Una capacidad

		longitud := len(meses)
		capacidad := cap(meses)

		fmt.Printf("La longitud es: %v, la capacidad es: %v %p\n", longitud, capacidad, meses)

		meses = append(meses, "Octubre") //Si se encuentra en capacidad máxima se genera un nuevo arreglo

		meses = append(meses, "Noviembre")
		meses = append(meses, "Diciembre")

		longitud = len(meses)
		capacidad = cap(meses)

		fmt.Printf("La longitud es: %v, la capacidad es: %v %p\n", longitud, capacidad, meses)
	*/
}
