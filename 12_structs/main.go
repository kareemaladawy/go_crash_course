package main

import (
	"fmt"
	"strconv"
)

// Struct = Class

// Define person struct
type Person struct {
	name, city, email string
	age               int

	// Two types of methods for a struct, value receivers & pointer receivers
	// Value receivers: When making calculations or not changing something
	// Pointer receivers: When changing something
}

// Define a value-receiver method on the Person struct
func (p Person) greet() string {
	return "Hello, my name is " + p.name + " and my age is " + strconv.Itoa(p.age)
}

// Define a pointer-receiver method on the Person struct
func (p *Person) incrAge() {
	p.age++
}

func (p *Person) getMarried() {
	p.name = p.name + " Jr."
}

func main() {
	// Init a person using the Person struct
	person1 := Person{
		name:  "Kareem",
		age:   25,
		city:  "New York",
		email: "email.com",
	}

	fmt.Println(person1.name, person1.age, person1.city, person1.email)
	fmt.Println(person1.greet())
	fmt.Println(person1.age)
	person1.incrAge()
	person1.getMarried()
	fmt.Println(person1.greet())
}
