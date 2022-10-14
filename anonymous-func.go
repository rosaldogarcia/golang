// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type people struct{}

func (p people) speak() {}

type human interface {
	speak()
}

func bar(h human) {}

func main() {

	// Anonymous function
	a := func(x int) {
		fmt.Println("The value of x: ", x)
	}
	a(46)
}

