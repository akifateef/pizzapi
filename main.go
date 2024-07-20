package main

import (
	"fmt"
	"math"
)

func main() {
	people, slicesPerPerson, slicesPerPizza := input()

	totalPizzasNeeded := calculatePizzaNeeded(slicesPerPerson, people, slicesPerPizza)

	output(totalPizzasNeeded)
}

func output(totalPizzasNeeded float64) (int, error) {
	return fmt.Print("You need to order ", totalPizzasNeeded, " Pizzas\n")
}

func input() (int, int, int) {
	var people int
	var slicesPerPerson int
	var slicesPerPizza int
	fmt.Print("Let's calculate how much Pizza you should get :)\n")

	fmt.Print("How many slices each pizza will have? ")
	fmt.Scan(&slicesPerPizza)

	fmt.Print("How many guests are there? ")
	fmt.Scan(&people)

	fmt.Print("How many slices each person would eat? ")
	fmt.Scan(&slicesPerPerson)

	fmt.Print("Number of slices per pizza: ", slicesPerPizza, "\n")
	fmt.Print("Number of people: ", people, "\n")
	fmt.Print("Number of slices per person: ", slicesPerPerson, "\n")
	
	return people, slicesPerPerson, slicesPerPizza
}

func calculatePizzaNeeded(slicesPerPerson int, people int, slicesPerPizza int) float64 {
	var totalSlicesNeeded int
	totalSlicesNeeded = slicesPerPerson * people

	return math.Ceil(float64(totalSlicesNeeded) / float64(slicesPerPizza))
}
