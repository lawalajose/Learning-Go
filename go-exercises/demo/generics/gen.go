package main

import "fmt"

type Leagues string
type Clubs string

func main() {
	leagues := make(map[Leagues][]Clubs)
	leagues["La liga"] = []Clubs{"Madrid", "Barca", "Atletico"}
	leagues["Ligue 1"] = []Clubs{"PSG", "Monaco", "PSV"}
	leagues["Premier League"] = []Clubs{"Chelsea", "Liverpool", "Arsenal"}
	leagues["La liga"] = []Clubs{"Elche", "Bibao", "Deportivo"}

	fmt.Println(leagues)

}
