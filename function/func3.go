// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	//unfurling slice

	ix := []int{1, 2, 3, 4, 5, 6, 78, 9, 10, 1}
	x := sum(ix...)
	fmt.Println("Total is: ", x)

}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, v, sum)
	}
	//fmt.Println("Total is: ", sum)
	return sum
}

