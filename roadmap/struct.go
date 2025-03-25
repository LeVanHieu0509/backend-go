package main

import "fmt"

type student struct {
	name string
	age  int
}
type person struct {
	name string
	age  int
	student
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func updatePerson(p *person, age int) {
	p.age = age
}

func main() {

	fmt.Println(person{"Bob", 20, student{"Alice", 20}})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50, student: student{"Alice", 20}}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	updatePerson(&s, 23)
	fmt.Println(sp.student.age)
}
