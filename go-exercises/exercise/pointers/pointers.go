//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import (
	"fmt"
)

type Store struct {
	Items       string
	SecurityTag bool
}

func actdeactTags(store *Store) {
	switch store.SecurityTag {
	case true:
		store.SecurityTag = false
	case false:
		store.SecurityTag = true
	}
}

func checkout(slices []Store) {
	for i := 0; i < len(slices); i++ {
		store := &slices[i]
		store.SecurityTag = false
	}
}

func printSlice(slice []Store) {
	for _, ch := range slice {
		fmt.Println(ch)
	}
}

func main() {
	//  - Create at least 4 items, all with active security tags
	item1 := Store{"Iphone", true}
	item2 := Store{"Laptop", true}
	item3 := Store{"Ipad", true}
	item4 := Store{"Iwatch", true}
	//  - Store them in a slice or array
	products := []Store{item1, item2, item3, item4}
	printSlice(products)
	fmt.Println()
	//  - Deactivate any one security tag in the array/slice
	actdeactTags(&products[1])
	printSlice(products)
	fmt.Println()

	// - Call the checkout() function to deactivate all tags
	checkout(products)
	printSlice(products)
	fmt.Println()
	//  - Print out the array/slice after each change

}
