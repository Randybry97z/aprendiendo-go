package multipleInterface

import "fmt"

//MULTIPLES INTERFACES
type Animal interface {
	Dormir()
}

type Mascota interface {
	Comer()
}

//ESTRUCTURAS
type Gato struct {
	Nombre string
}

func (self Gato) Dormir() {
	fmt.Println("El gato duerme")
}

func (self Gato) Comer() {
	fmt.Println("El gato come")
}

func funcionParaAnimal(animal Animal) {
	fmt.Println("El objeto es un animal")
}

func funcionParaMascota(mascota Mascota) {
	fmt.Println("El objeto es una mascota")
}

func multipleInterface() {

	gato := Gato{Nombre: "Mike"}

	gato.Dormir()
	gato.Comer()

	funcionParaAnimal(gato)
	funcionParaMascota(gato)

}
