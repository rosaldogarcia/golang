// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// struct
type people struct {
	first string
	last  string
}

// method
func (p people) speak() {
	fmt.Println(p.first, p.last)
}

// interface
type human interface {
	speak()
}

func bar(h human) {

	switch h.(type) {
	case people:
		fmt.Println("type people")
	default:
		fmt.Println("default")
	}

}
func main() {

	p1 := people{
		first: "James",
		last:  "Bond",
	}
	p1.speak()
	bar(p1)

	//map
	fmt.Println("------ MAP --------")

	m := map[string]int{
		"Diana": 3,
	}
	fmt.Println(m["Diana"])
	v, ok := m["Diana"]
	fmt.Println("Value of v: ", v)
	fmt.Println("Value of ok: ", ok)
	if v, ok := m["Diana"]; ok {
		fmt.Println("This is the IF value of MAP : ", v)
	}

	//make
	fmt.Println("------ MAKE --------")
	z := make([]int, 3, 3)
	fmt.Println("This is make output: ", z)

	//slice
	fmt.Println("------ SLICE --------")
	w := []string{}
	fmt.Println(w)

	// Function (blank, passing value)
	fmt.Println("------ FUNCTION --------")

	fi := func() {
		fmt.Println("This is the blank function")
	}
	fi()

	f := func(x int) {
		fmt.Println("This is the value of x", x)
	}
	f(46)
}

