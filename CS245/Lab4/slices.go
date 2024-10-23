package main

import "fmt"

type colors struct {
	number int
	italian string
	decimal float32
}

func (item colors) String() string {
	return fmt.Sprintf("Color number:%v  italian:%v  decimal:%v", item.number, item.italian, item.decimal);
}

type slices []colors

func main() {
	pink := colors{8, "rosa", 8.8}
	orange := colors{2, "arancia", 2.2}
	indigo := colors{6, "indaco", 6.6}
	green := colors{4, "verde", 4.4}
	red := colors{1, "rossa", 1.1}
	blue := colors{5, "blu", 5.5}
	violet := colors{7, "viola", 7.7}
	yellow := colors{3, "gialla", 3.3}
	teal := colors{10, "verde acqua", 10.10}
	vermilion := colors{9, "vermiglio", 9.9}

	rainbow := slices{red, orange, yellow, green, blue, indigo, violet, pink, vermilion, teal}

	fmt.Println(rainbow)
}