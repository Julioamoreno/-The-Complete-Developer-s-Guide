package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	user := person{
		firstName: "Julio",
		lastName:  "Armando",
		contactInfo: contactInfo{
			email:   "julioarmando321@gmail.com",
			zipCode: 4500000,
		},
	}

	user.updateName("test")
	user.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
