// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`
	
	[{"First":"Diana","Last":"Kulet","Age":4},{"First":"Chloe","Last":"Pachochi","Age":10},{"First":"Leanne","Last":"Bujingjing","Age":16}]

	`)

	type Person struct {
		First string
		Last  string
		Age   int
	}

	var person []Person

	err := json.Unmarshal(jsonBlob, &person)

	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(person)
}

