//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	Name  string
	Price float32
}

func printItem(shoppingList [4]Product) {
	for i := 0; i < len(shoppingList); i++ {
		items := shoppingList[i]
		fmt.Println(items)
	}
}

func main() {
	shoppingList := [4]Product{
		{"Iphone 17 Pro Max", 45.07},
		{"MacBook Pro", 23.21},
		{"Nike AirMax", 5.67},
	}
	printItem(shoppingList)

	fmt.Println(shoppingList)
	//* Print to the terminal:
	//  - The last item on the list
	lastItem := shoppingList[len(shoppingList)-1]
	fmt.Println(lastItem)
	// - The total number of items
	numberOfItems := len(shoppingList)
	fmt.Println(numberOfItems)
	// - The total cost of the items
	totalCost := func() float32 {
		var sum float32
		for i := 0; i < len(shoppingList); i++ {
			sum += shoppingList[i].Price
		}
		return sum
	}
	fmt.Println(totalCost())
	//* Add a fourth product to the list
	shoppingList[len(shoppingList)-1] = Product{"LCD Screen TV", 70.34}
	//  and print out the
	//  information again
	printItem(shoppingList)

	number := []int{1, 2, 3}
	num := []int{4, 6, 8}
	green := num[:]
	addnumber := append(num, number...)
	fmt.Println(addnumber)

	// animal := []string{"goat", "cat", "dog"}

}
