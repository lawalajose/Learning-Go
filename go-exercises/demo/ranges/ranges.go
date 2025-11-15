package main

import "fmt"

func main() {
	intro := []string{"Hello", "World", "!"}
	for i := 0; i < len(intro); i++ {

		fmt.Println(i, intro[i])

		for _, r := range intro[i] {
			fmt.Printf("  %q\n", r)
		}
	}
}
