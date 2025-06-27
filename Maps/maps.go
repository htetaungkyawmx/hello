package main

import "fmt"

func main() {

	/*StudentAge := make(map[string]int)

	StudentAge["James"] = 21
	StudentAge["John"] = 22
	StudentAge["Jonny"] = 23
	StudentAge["Jose"] = 24
	StudentAge["Jin"] = 25
	StudentAge["Jak"] = 26*/

	/*fmt.Println(StudentAge["James"])*/

	/*fmt.Println(len(StudentAge))*/

	// Create a nested map:
	// Outer map: superhero name → Inner map (with RealName and City)
	superhero := map[string]map[string]string{

		"Superman": map[string]string{
			"RealName": "Clark Kent",
			"City":     "San Francisco",
		},

		"Batman": map[string]string{
			"RealName": "Bruce Wayne",
			"City":     "Gotham",
		},
	}

	// Check if "Superman" exists in map
	if temp, hero := superhero["Batman"]; hero {
		// If found, print RealName and City
		fmt.Println(temp["RealName"], temp["City"]) // Clark Kent San Francisco
	}
}
