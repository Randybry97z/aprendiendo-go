package variadic

//VARIADIV FUNCTION (permite n parametros en la función)

/* //VARIABLES Y SCOPES
func modificarVariable(param string) {
	fmt.Println("Valor actual: ", param)
	param = "Cambio de valor"
	fmt.Println("Nuevo valor: ", param)
} */

/* //VARIADIC FUNCTIONS
func promedio(calificaciones ...int) int {

	var promedio, suma int
	for _, calificacion := range calificaciones {
		suma = suma + calificacion
	}
	promedio = suma / len(calificaciones)
	return promedio
} */

func variadic() {
	/* 	//variables y scopes
		var curso = "Curso Go"

		modificarVariable(curso)

	   	fmt.Println(">>>", curso) */

	/* //VARIADIC FUNCTIONS
	resultado := promedio(10, 6, 8, 8, 10, 7, 9, 10)

	fmt.Println("Tu promedio es: ", resultado) */
}
