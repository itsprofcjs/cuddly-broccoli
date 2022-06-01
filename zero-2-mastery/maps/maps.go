package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(serverStatus map[string]int) {
	var noOfServer = len(serverStatus)
	statusCount := make(map[int]int)

	for _, status := range serverStatus {
		statusCount[status] += 1
	}

	fmt.Println("No of Server", noOfServer)
	fmt.Println("Status count", statusCount)

}

func main() {
	shoppingList := make(map[string]int)

	fmt.Println(len(shoppingList))

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 6

	shoppingList["eggs"] += 5

	fmt.Println(shoppingList)

	delete(shoppingList, "milk")

	fmt.Println("Milk deleted, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]

	fmt.Println("Need cereal?")

	if !found {
		fmt.Println("Nope")
	} else {
		fmt.Println("found", cereal)
	}

	fmt.Println(len(shoppingList))

	totalItems := 0

	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("There are", totalItems, "on the shopping list")

	//--Summary:
	//  Write a program to display server status.
	//
	//--Requirements:
	//* Create a function to print server status, including:
	//  - Number of servers
	//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
	//* Store the existing slice of servers in a map
	//* Default all of the servers to `Online`
	//* Perform the following status changes and display server info:
	//  - display server info
	//  - change `darkstar` to `Retired`
	//  - change `aiur` to `Offline`
	//  - display server info
	//  - change all servers to `Maintenance`
	//  - display server info

	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatus := make(map[string]int)

	for _, server := range servers {
		serverStatus[server] = Online
	}

	printServerStatus(serverStatus)

	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline

	delete(serverStatus, "baseline")

	printServerStatus(serverStatus)

	for key := range serverStatus {
		serverStatus[key] = Maintenance
	}

	printServerStatus(serverStatus)
}
