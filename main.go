package main

import "fmt"

type Cars struct {
	Model string
	Year  int
}

func main() {
	cars := []Cars{
		{Model: "Fiat Honda", Year: 2022},
		{Model: "Maverick", Year: 2000},
		{Model: "Porsche 911", Year: 2023},
	}
	var ordered []Cars
	for range cars {
		higher := higherSearch(cars)
		ordered = append(ordered, cars[higher])
		cars = append(cars[:higher], cars[higher+1:]...)
	}

	fmt.Println("AA ", ordered)
}
func higherSearch(cars []Cars) int {
	higher := 0
	for k, v := range cars {
		if v.Year > cars[higher].Year {
			higher = k
		}
	}
	return higher
}