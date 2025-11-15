package main

import "fmt"

func main() {

	myMap := map[string]int{
		"Barcelona":       20,
		"Real Madrid":     19,
		"Real Sociedad":   18,
		"Atletico Bilbao": 18,
	}
	myMap2 := make(map[string]int)
	myMap2["Arsenal"] = 35
	myMap2["Sunderland"] = 32
	myMap2["Manchester City"] = 30

	fmt.Println(myMap)
	fmt.Println(myMap2)

	delete(myMap, "Atletico Bilbao")
	myMap2["Girona"] = 14
	fmt.Println(myMap)
	fmt.Println(myMap2)

	barca, found := myMap["Barcelona"]
	fmt.Println(barca)
	fmt.Println(found)

	fmt.Println(myMap)

	for k, v := range myMap {
		fmt.Println(k, v)
	}

}
