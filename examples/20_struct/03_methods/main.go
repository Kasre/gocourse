package main

type Person struct {
	name string
}

func (p *Person) speak() string {
	return "Hello!"
}

type Programmer struct {
	Person
}

func (p *Programmer) speak() string {
	return "42 is The Answer to the Ultimate Question of Life"
}

func main() {
	p := Programmer{Person{"Moses"}}

	println(p.Person.speak())
}
