package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// variable declarations
	var color string = "red"
	age := 20

	// function and return
	show(color)
	show(age)

	show(calculate(1, 2))

	// array & slice
	// slice
	var colors []string
	colors = append(colors, "red")
	show(colors)

	var cars = []string{"ferrari", "bmw"}
	show(cars)

	// array
	var countries = [2]string{"ID", "US"}
	show(countries)

}

func show(i interface{}) {
	fmt.Println(i)
}

func calculate(a int, b int) int {
	return a + b
}
