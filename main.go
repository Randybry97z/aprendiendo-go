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

		//MAKE
		slice := make([]int, 3, 3)

		slice[0] = 100
		slice[1] = 200
		slice[2] = 300

		slice = append(slice, 400)

		fmt.Println(slice)
		fmt.Println(len(slice))
		fmt.Println(cap(slice))
	*/

	/*
		//MAPS (Collecion desordenada de elementos)

		dias := make(map[int]string)

		dias[0] = "Domingo"
		dias[1] = "Lunes"
		dias[2] = "Martes"
		dias[3] = "Miercoles"
		dias[4] = "Jueves"
		dias[5] = "Viernes"
		dias[6] = "Sabado"

		dias[4] = "JUEVES"

		///Eliminar llave
		delete(dias, 4)

		fmt.Println(dias)

		usuarios := make(map[string][]int)
		usuarios["Bryan"] = []int{10, 9, 10, 7, 5}

		fmt.Println(usuarios)

		//ITERACION DE MAPAS

		usuarios := map[int]string{} //Se crea variable tipo mapa

		usuarios[1] = "Usuario 1"
		usuarios[2] = "Usuario 2"
		usuarios[3] = "Usuario 3"

		for id, username := range usuarios {
			fmt.Println(id, username)
		}
	*/
}
