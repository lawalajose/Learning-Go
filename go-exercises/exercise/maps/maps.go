//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServers(myMap map[string]int) {
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	status := make(map[int]int)
	for _, server := range myMap {
		switch server {
		case Online:
			status[Online] += 1
		case Offline:
			status[Offline] += 1
		case Maintenance:
			status[Maintenance] += 1
		case Retired:
			status[Retired] += 1
		default:
			panic("Unhandled server")
		}
	}

	for key, value := range status {
		fmt.Println(key, value)
	}

	fmt.Println(status[Online], "people are Online")
	fmt.Println(status[Offline], "people are Offline")
	fmt.Println(status[Maintenance], "people are Maintainance")
	fmt.Println(status[Retired], "people are Retired")

}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	myMap := make(map[string]int)
	for _, r := range servers {
		myMap[r] = Online
	}
	displayServers(myMap)
	fmt.Println()
	//  - change server status of `darkstar` to `Retired`
	myMap["darkstar"] = Retired

	// - change server status of `aiur` to `Offline`
	myMap["aiur"] = Offline
	// - call display server info function
	displayServers(myMap)
	fmt.Println()
	// - change server status of all servers to `Maintenance`
	for key, _ := range myMap {
		myMap[key] = Maintenance
	}
	// - call display server info function
	displayServers(myMap)

}
