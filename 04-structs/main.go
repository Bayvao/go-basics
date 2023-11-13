package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	
	// first way to use a struct 
	bruce := person{"Bruce", "Wayne", contactInfo {"bruce@gmail.com", 1234}}
	fmt.Println(bruce)
	
	// second way to use a struct 
	clark := person{firstName: "Clark", lastName: "Kent"}
	fmt.Println(clark)

	// third way to use a struct 
	var barry person
	barry.firstName = "Barry"
	barry.lastName = "Allen"
	fmt.Println(barry)
	
	/*
	* use a nested struct
	*/
	peter := person{
		firstName: "Peter",
		lastName:  "Parker",
		//contact: contactInfo{
		contactInfo: contactInfo {
			email:   "peter@gmail.com",
			zipCode: 9400,
		},
	}

	peter.updateFirstName("Pater")
	peter.display()
}

func (pointerToPerson *person) updateFirstName (newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) display() {
	fmt.Printf("%+v", p)
}