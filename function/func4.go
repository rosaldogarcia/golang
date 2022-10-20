// You can edit this code!
// Click here and start typing.
// https://go.dev/play/p/TM8MCtJPRDC

package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("%T\n", x)
}

func foo() string {
	s := "Hello,World"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}

