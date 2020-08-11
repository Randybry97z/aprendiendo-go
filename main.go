package main

import "fmt"

//Exec go build (compila y genera ejecutable) || go run compila internamente

func main() {
	var name string // default ==  ""
	var age int     // default == 0
	var lastName string = "Sandoval"

	name = "Bryan Saúl"
	age = 23

	city := "México"
	postalCode := 55635

	var height = 1.73

	fmt.Println(name, lastName)
	fmt.Println(age, height)
	fmt.Println(city, postalCode)
}
