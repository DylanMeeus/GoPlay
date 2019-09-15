package main

import "fmt"

func main() {
	p := GeneratePerson("dylan")
	p.IfPresent(func(pp Person) { fmt.Println(pp) })
}

func GeneratePerson(name string) PersonMonad {
	if name == "" {
		return NewPersonMonad(nil)
	} else {
		return NewPersonMonad(&Person{"Dylan", "Meeus"})
	}
	return PersonMonad{}
}

type Person struct {
	Name, LastName string
}

type PersonMonad struct {
	Value   Person
	Present bool
}

func (i PersonMonad) Get() Person {
	return i.Value
}

func (i PersonMonad) IsPresent() bool {
	return i.Present
}

func (i PersonMonad) IfPresent(f func(i Person)) {
	if i.Present {
		f(i.Value)
	}
}

func (i PersonMonad) OrElse(e Person) Person {
	if !i.Present {
		return e
	}
	return i.Value
}

func NewPersonMonad(p *Person) PersonMonad {
	if p == nil {
		return PersonMonad{Person{}, false}
	}
	return PersonMonad{*p, true}
}
