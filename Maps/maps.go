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

	superhero := map[string]map[string]string{

		"Superman": map[string]string{
			"RealName": "Clark Kent",
			"City":     "San Francisco",
		},

		"Batman": map[string]string{
			"RealName": "Bruce Wayne",
			"City":     "San Francisco",
		},
	}

	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}
