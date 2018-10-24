package main

import (
	"fmt"
	"strings"

	"../pack"
	"../pack1"
)

//Person 是一个人
type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
func main() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)

	str := ReturnStr()
	fmt.Println(str)

	str1 := pack.Sys()
	fmt.Println(str1)

	var pers1 Person
	pers1.firstName = "min"
	pers1.lastName = "li"
	upPerson(&pers1)
	fmt.Printf("the name of the person is %s %s \n", pers1.firstName, pers1.lastName)

	string1 := pack1.ReturnString()
	fmt.Println(string1)

}
