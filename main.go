package main

import (
	"fmt"
	"math"
)

func main() {
	var people int
	var slicesPerPerson int
	var slicesPerPizza int
	fmt.Print("Let's calculate how much Pizza you should get :)\n")

	fmt.Print("How many guests are there? ")
	fmt.Scan(&people)

	fmt.Print("How many slices each person would eat? ")
	fmt.Scan(&slicesPerPerson)

	fmt.Print("How many slices each pizza will have? ")
	fmt.Scan(&slicesPerPizza)

	fmt.Print("Number of people: ", people, "\n")
	fmt.Print("Number of slices per person: ", slicesPerPerson, "\n")
	fmt.Print("Number of slices per pizza: ", slicesPerPizza, "\n")

	var totalSlicesNeeded int
	totalSlicesNeeded = slicesPerPerson * people

	var totalPizzasNeeded float64
	totalPizzasNeeded = math.Ceil(float64(totalSlicesNeeded) / float64(slicesPerPizza))

	fmt.Print("You need to order ", totalPizzasNeeded, " Pizzas\n")
}
