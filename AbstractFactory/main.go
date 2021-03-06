package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")
	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort(10)
	adidasShoe := nikeFactory.makeShoe()
	adidasShort := adidasFactory.makeShort(20)

	printShoeDetails(nikeShoe)
	printShoeDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShoeDetails(adidasShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShortDetails(s iShort) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
