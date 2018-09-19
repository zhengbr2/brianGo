package main

import (
	"fmt"
	"unsafe"
)

func main2() {

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

	pers := []string{}
	fmt.Println("sizeof:", unsafe.Sizeof(pers))
	fmt.Println("pers:", pers, "has new length", len(pers), "and capacity", cap(pers))
	pers = append(pers, "brian", "quan", "fey")
	fmt.Println("pers:", pers, "has new length", len(pers), "and capacity", cap(pers))

}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
func main() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
