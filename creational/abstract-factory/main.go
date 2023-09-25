package main

import "fmt"

func main() {
	nikeFactory, nikeErr := GetSportsFactory("nike")
	if(nikeErr != nil) {
		fmt.Println(nikeErr)
	}

	adidasFactory, adidasErr := GetSportsFactory("adidas")
	if(adidasErr != nil) {
		fmt.Println(adidasErr)
	}

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	fmt.Println("Abstract Factory test:")
	fmt.Println("======================")

	printDetails(nikeShoe, nikeShirt)
	printDetails(adidasShoe, adidasShirt)

}

func printDetails(shoe IShoe, shirt IShirt) {
	fmt.Println("Shoe Details: ")
	fmt.Println("Logo: ", shoe.getLogo())
	fmt.Println("Size: ", shoe.getSize())
	fmt.Println(" ")
	fmt.Println("Shirt Details: ")
	fmt.Println("Logo: ", shirt.getLogo())
	fmt.Println("Size: ", shirt.getSize())
	fmt.Println(" ")
}