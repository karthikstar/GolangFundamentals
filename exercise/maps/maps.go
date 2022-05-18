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

// server statuses
const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(m map[string]int) {
	serverCount := 0
	OnCount, OffCount, MCount, RCount := 0, 0, 0, 0
	for _, status := range m {
		serverCount += 1
		switch status {
		case Online:
			OnCount += 1
		case Offline:
			OffCount += 1
		case Maintenance:
			MCount += 1
		case Retired:
			RCount += 1
		default:
			panic("unhandled server status")
		}
	}
	fmt.Println("Total no of servers", serverCount)
	fmt.Println("Offline", OffCount)
	fmt.Println("Online", OnCount)
	fmt.Println("Maintenance", MCount)
	fmt.Println("Retired", RCount)
}

func main() {
	//servers
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverMap := make(map[string]int) // key will be the server name, value will be the erver status
	for i := 0; i < len(servers); i++ {
		serverMap[servers[i]] = Online
	}
	//display server info
	printServerStatus(serverMap)

	// change darkstar server to retired
	serverMap["darkstar"] = Retired
	// make aiur server offline
	serverMap["aiur"] = Offline
	printServerStatus(serverMap)
	for key := range serverMap {
		serverMap[key] = Maintenance
	}
	printServerStatus(serverMap)

}
