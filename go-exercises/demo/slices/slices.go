package main

import "fmt"

func printSlice(route string, slice []string) {
	fmt.Println()
	fmt.Println("---", route, "---")
	for i := 0; i < len(slice); i++ {
		routes := slice[i]
		fmt.Println(routes)
	}
}

func main() {
	routes := []string{"Office", "Library", "Supermarket"}
	printSlice("Route 1", routes)

	routes = append(routes, "Gym")
	printSlice("Route 2", routes)
	fmt.Println()
	fmt.Println()

	fmt.Println(routes[0], "visited")
	fmt.Println(routes[1], "visited")

	routes = routes[2:]
	printSlice("Remaining locations", routes)

}
