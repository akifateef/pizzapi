package main

import (
	"fmt"
	"math"
)

func main() {
	var numberOfPeople int
	var numberOfSlicesPerPizza int
	var numberOfSlicesPerPerson int
	fmt.Print("Let's calculate how much Pizza you should get :)\n")

	fmt.Print("How many guests are there? ")
	fmt.Scan(&numberOfPeople)

	fmt.Print("How many slices each person would eat? ")
	fmt.Scan(&numberOfSlicesPerPerson)

	fmt.Print("How many slices each pizza will have? ")
	fmt.Scan(&numberOfSlicesPerPizza)

	fmt.Print("number_of_people: ", numberOfPeople, "\n")
	fmt.Print("number_of_slices_per_person: ", numberOfSlicesPerPerson, "\n")
	fmt.Print("number_of_slices_per_pizza: ", numberOfSlicesPerPizza, "\n")

	var totalNumberOfSlicesNeeded int
	totalNumberOfSlicesNeeded = numberOfSlicesPerPerson * numberOfPeople

	var totalNumberOfPizzasNeeded float64
	totalNumberOfPizzasNeeded = math.Ceil(float64(totalNumberOfSlicesNeeded / numberOfSlicesPerPizza))

	fmt.Print("You need to order ", totalNumberOfPizzasNeeded, " Pizzas\n")
}
