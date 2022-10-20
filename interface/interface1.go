///https://go.dev/play/p/Qu1XN2UoKJp

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type people struct {
	first string
	last  string
}

type secretAgent struct {
	people
	ltk bool
}

func (p people) speak() {
	fmt.Println(p.first, p.last)
}

func (s secretAgent) speak() {
	fmt.Println(s.first, s.last, s.ltk)
}

type human interface {
	speak()
}

func bar(h human) {

	switch h.(type) {
	case people:
		fmt.Println("Person Type")
	case secretAgent:
		fmt.Println("SecretAgent Type")
	default:
		fmt.Println("Default.")
	}

	fmt.Println("I was pass into bar I called human", h)
}

func main() {

	per1 := people{
		first: "Dr.",
		last:  "Yes",
	}

	sa1 := secretAgent{
		people: people{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
	sa1.speak()
	per1.speak()
	bar(sa1)
	bar(per1)
}

