package main

import "fmt"

func main() {
	e := Employee{Person: Person{
		Name: "张",
		Age:  18,
	},
		EmployeeID: 50}
	e.info()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (p Employee) info() {
	fmt.Printf("id=%d,name=%s, age=%d", p.EmployeeID, p.Name, p.Age)
}
