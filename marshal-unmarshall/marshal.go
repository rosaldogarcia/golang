// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		First string
		Last  string
		Age   int
	}
	p1 := Person{"Diana", "Kulet", 4}
	p2 := Person{"Chloe", "Pachochi", 10}
	p3 := Person{"Leanne", "Bujingjing", 16}

	person := []Person{p1, p2, p3}

	bs, err := json.Marshal(&person)

	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(bs))
}

