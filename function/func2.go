// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 78, 9, 10, 1)
	fmt.Println("Total is: ", x)

	fmt.Println(mouse("Diana", "Amaris"))

	a, b := wolverine("Logan", "Howlett")
	fmt.Println(a)
	fmt.Println(b)

}

func wolverine(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln)
	b := false
	return a, b
}

func mouse(fn string, ln string) string {
	return fmt.Sprint(fn, ln)
}

func sum(x ...int) int {
	// fmt.Println(x)
	// fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
		//fmt.Println(i, v, sum)
	}
	// fmt.Println("Total is: ", sum)
	return sum
}

