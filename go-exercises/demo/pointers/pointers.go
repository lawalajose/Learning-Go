package main

import "fmt"

type Counter struct {
	hits int
}

func increment(Counter *Counter) {
	Counter.hits += 1
}

func converter(old *string, new string, Counter *Counter) {
	*old = new
	increment(Counter)
}

func main() {

	hello := "Hello"
	world := "World"

	Counter := Counter{}
	fmt.Println(hello, world, Counter)
	fmt.Println()

	converter(&hello, "Hi", &Counter)
	fmt.Println(hello, world, Counter)
	fmt.Println()

	converter(&world, "Go!", &Counter)
	fmt.Println(hello, world, Counter)
	fmt.Println()

}
