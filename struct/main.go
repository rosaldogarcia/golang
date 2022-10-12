// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type truck struct {
	doors     int
	color     string
	fourWheel bool
	contacts  map[string]int
}
type sedan struct {
	doors     int
	color     string
	luxury    bool
	passenger []string
}

type vehicle struct {
	truck
	sedan
	makers []string
}

func main() {

	t1 := truck{
		doors:     2,
		color:     "White",
		fourWheel: false,
		contacts: map[string]int{
			"Diana":  1111,
			"Chloe":  222,
			"Bujing": 333,
		},
	}
	fmt.Println(t1.doors, t1.color, t1.fourWheel)
	for i, v := range t1.contacts {
		fmt.Println(i, v)
	}

	c1 := sedan{
		doors:  4,
		color:  "Red",
		luxury: true,
		passenger: []string{
			"Diana",
			"Chloe",
			"Bujig",
		},
	}
	fmt.Println(c1.doors, c1.color, c1.luxury)
	for _, k := range c1.passenger {
		fmt.Println(k)
	}

	v1 := vehicle{
		truck: truck{
			doors:     2,
			color:     "Black",
			fourWheel: false,
			contacts: map[string]int{
				"Diana":  111,
				"Chloe":  222,
				"Bujing": 333,
			},
		},
		sedan: sedan{
			doors:  4,
			color:  "Black",
			luxury: false,
			passenger: []string{
				"Diana",
				"Chloe",
				"Bujing",
			},
		},
		makers: []string{
			"Ferrari",
			"Lamborgini",
			"Jaguar",
			"Mercedez",
		},
	}
	fmt.Println(v1.truck.doors, v1.truck.color, v1.truck.fourWheel)
	for l, m := range v1.truck.contacts {
		fmt.Println(l, m)
	}
	fmt.Println(v1.sedan.doors, v1.sedan.color, v1.sedan.luxury)
	for n, o := range v1.sedan.passenger {
		fmt.Println(n, o)
	}
	for p, q := range v1.makers {
		fmt.Println(p, q)
	}

}

