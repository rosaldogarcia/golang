// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//////////
// struct
/////////

type people struct {
	first string
	last  string
}

type secretAgent struct {
	people
	ltk bool
}

//////////
// method
/////////

func (p people) speak() {
	fmt.Println(p.first, p.last)
}

func (s secretAgent) speak() {
	fmt.Println(s.first, s.last)
}

//////////////
// interface
/////////////

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I pass into human.", h)
	switch h.(type) {
	case people:
		fmt.Println("type person.")
	case secretAgent:
		fmt.Println("type secretAgent.")
	default:
		fmt.Println("default")
	}
}

func main() {

	p1 := people{
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

	// struct
	fmt.Println(p1.first, p1.last)
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
	// method
	p1.speak()
	sa1.speak()
	bar(p1)
	bar(sa1)

}

