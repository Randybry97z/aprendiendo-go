package metodosObjeto

import "fmt"

type User struct {
	Name  string
	Email string
}

func (self *User) setName(name string) {
	self.Name = name
}

func (self *User) getName() string {
	return self.Name
}

func metodosObjeto() {

	cody := User{Name: "Cody", Email: "cody@cody.com"}

	cody.setName("Bryan")

	fmt.Println(cody.getName())

}
