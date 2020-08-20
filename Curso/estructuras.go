package estructuras

import "fmt"

//CREAR ESTRUCTURAS
type User struct {
	Name  string //""
	Email string //""
	Age   int    //0
}

func estructuras() {

	//ESTRUCTURA CREADA
	//var cody User //Un objeto

	//cody.Name = "Cody"
	//cody.Email = "cody@cody.com"
	//cody.Age = 20

	Name := "Cody"
	Email := "cody@cody.com"
	Age := 20
	cody := User{Name, Email, Age} //Definición del objeto

	fmt.Println(cody)
}
