// https://go.dev/play/p/EppB4ZoV7bx

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

func (s people) speak() {
	fmt.Println(s.first, s.last)
}
func main() {

	p1 := people{
		first: "Diana",
		last:  "Amaris",
	}
	p2 := people{
		first: "Chloe",
		last:  "Anne",
	}

	sa1 := secretAgent{
		people: people{
			first: "Chloe",
			last:  "Anne",
		},
		ltk: true,
	}
	fmt.Println(p1.first, p1.last)
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
	p2.speak()
	p1.speak()
}

