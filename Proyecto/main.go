package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var id int
var users map[int]User

func crearUsuario() {
	clearConsole()
	fmt.Println("Ingresa un valor para username: ")
	username := readLine()

	fmt.Println("Ingresa un valor para email: ")
	email := readLine()

	fmt.Println("Ingresa un valor para edad: ")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir de un string a un entero")
	}

	id++

	user := User{id, username, email, age}
	users[id] = user

	//fmt.Println(users)

	fmt.Println(">>> Usuario creado")
}

func listarUsuarios() {
	clearConsole()
	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}

	fmt.Println("\n")
}

func actualizarUsuario() {
}

func eliminarUsuario() {
	clearConsole()
	fmt.Print("Ingresa el id del usuario a eliminar: ")
	id, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir string a entero")
	}

	if _, ok := users[id]; ok {
		delete(users, id)
	}

	fmt.Println(">>> Usuario eliminado\n")
}

func clearConsole() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string {

	if option, err := reader.ReadString('\n'); err != nil {
		panic("no es posible obtener el valor!")
	} else {
		cadena := strings.TrimSpace(option)

		return cadena
	}

}

func main() {

	users = make(map[int]User)
	reader = bufio.NewReader(os.Stdin)

	for {

		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")

		fmt.Print("Ingresa una opción valida: ")
		option := readLine()

		if option == "quit" || option == "q" {
			break
		}

		switch option {
		case "a", "crear":
			crearUsuario()
		case "b", "listar":
			listarUsuarios()
		case "c", "actualizar":
			actualizarUsuario()
		case "d", "eliminar":
			eliminarUsuario()
		default:
			fmt.Println("Opción no valida")
		}
	}

	fmt.Println("Adiós")

}
