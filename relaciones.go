package relaciones

import "fmt"

//RELACIONES

/* //1:1
type User struct {
	Name   string
	Email  string
	Active bool
}

type Student struct {
	User User
	Id   string
} */

//1:many
type Curso struct {
	Title  string
	Videos []Video
}

type Video struct {
	Title string
	Curso Curso
}

func relaciones() {

	/* //1:1 Relationship
	bryan := User{Name: "bryan", Email: "bryan@bryan.com", Active: true}

	saul := User{Name: "saúl", Email: "saul@bryan.com", Active: true}

	bryanStudent := Student{User: bryan, Id: "1f12nd"}

	fmt.Println(bryan)
	fmt.Println(saul)
	fmt.Println(bryanStudent) */

	//1:many Relationship
	video1 := Video{Title: "1. Introducción"}
	video2 := Video{Title: "2. Instalación"}
	video3 := Video{Title: "3. Variables"}

	videos := []Video{video1, video2, video3}

	curso := Curso{Title: "Curso de Go!", Videos: videos}

	video1.Curso = curso
	video2.Curso = curso
	video3.Curso = curso

	fmt.Println(video1.Curso.Title)

	for _, video := range curso.Videos {
		fmt.Println(video.Title)
	}
}
