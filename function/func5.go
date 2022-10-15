//https://go.dev/play/p/8voje4Kkw6C

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	s := foo()
	fmt.Println(s)

	fmt.Println(bar()())
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

