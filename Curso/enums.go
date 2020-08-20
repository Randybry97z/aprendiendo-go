package enums

import "fmt"

//ENUMS
type UserType int

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	Username string
	Type     UserType
}

func enums() {

	//ENUMS
	bryan := User{Username: "randy", Type: Student}
	saul := User{Username: "saúl", Type: Teacher}

	fmt.Println(bryan)
	fmt.Println(saul)

	if bryan.Type == Student {
		fmt.Println("El usuario", bryan.Username, "es estudiante")
	}

}
