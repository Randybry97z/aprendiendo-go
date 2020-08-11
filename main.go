package main

//Exec go build (compila y genera ejecutable) || go run compila internamente

//CONSTANTES

//const COURSE string = "Curso profesional de Go!"

func main() {

	//DECLARACIÓN DE VARIABLES
	/*
		var name string // default ==  ""
		var age int     // default == 0
		var lastName string = "Sandoval"

		name = "Bryan Saúl"
		age = 23

		city := "México"
		postalCode := 55635

		var height = 1.73

		fmt.Println("Name: ", name, lastName)
		fmt.Println("Age and height:", age, height)
		fmt.Println("Address: ", city, postalCode) */

	/*
		var name string
		var lastName string
		var country string
	*/

	//DECLARACIÓN DE MÚLTIPLES VARIABLES
	/* //var name, lastName, country string
	//var name, lastName, country = "Bryan", "Sandoval", "México"
	name, lastName, country := "Bryan", "Sandoval", "México"

	fmt.Println(name, lastName, country) */
	//fmt.Println(COURSE)

	/*
			OPERADORES RELACIONALES
			>
			<
			>=
			<=
			==
			!=

		var edad = 10
		var resultado = edad > 5

		fmt.Println(resultado)
	*/

	/* //OPERADORES LÓGICOS (&& ; || ; !)

	//resultado := 20 == 20 && 30 == 30
	//resultado := 20 > 40 && 30 == 30
	//resultado := 20 == 20 || 30 == 30 || 20 > 40
	resultado:= 15 == 15 && 60 < 100 && (90 < 100 || 100 == 90)
	fmt.Println(resultado) */
}
