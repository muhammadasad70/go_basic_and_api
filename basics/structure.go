package main

import (
	"fmt"
)

type ContactInfo struct {
	Age     int
	Email   string
	ZipCode int
}
type Person struct {
	FirstName string
	LastName  string
	contact   ContactInfo
}

func ForStructure() {
	person := Person{
		FirstName: getUserStringData("Please enter your First Name :"),
		LastName:  getUserStringData("Please enter your Last Name :"),
		contact: ContactInfo{
			Age:     getUserIntData("Enter your Age :"),
			Email:   getUserStringData("Enter your Email :"),
			ZipCode: getUserIntData("Enter your Zipcode :"),
		},
	}
	person.printValue()
	// UpdateThePersonInfo(&person)
	person.UpdateInfo()
	person.printValue()
	// for get the struct with the declear name we use the formate +v(+v for entire values in the struct )with printf
	// fmt.Printf("Details of the user is : %+v\n", person)
	// for one by one value
	// fmt.Printf("First Name is : %v and Last Name is : %v\n", person.FirstName, person.LastName)
}
func UpdateThePersonInfo(p *Person) {
	p.LastName = getUserStringData("Please enter your First Name :")

}

// this is method use with the structure
func (p *Person) UpdateInfo() {
	p.FirstName = getUserStringData("Please enter your First Name :")
	p.LastName = getUserStringData("Please enter your Last Name :")
	p.contact.Age = getUserIntData("Enter your age :")
	p.contact.Email = getUserStringData("Enter your email :")
	p.contact.ZipCode = getUserIntData("Enter youer city zipcode :")

}

// this is function of the structure mean we can also make the function of the structure
// we can also validate the user details by checking
func Newperson() Person {
	return Person{
		FirstName: getUserStringData("Please enter your First Name :"),
		LastName:  getUserStringData("Please enter your Last Name :"),
		contact: ContactInfo{
			Age:     getUserIntData("Enter your Age :"),
			Email:   getUserStringData("Enter your Email :"),
			ZipCode: getUserIntData("Enter your Zipcode :"),
		},
	}

}
func getUserStringData(text string) string {
	fmt.Println(text)
	var value string
	fmt.Scan(&value)
	return value
}
func getUserIntData(text string) int {
	fmt.Println(text)
	var value int
	fmt.Scan(&value)
	return value
}
func (p Person) printValue() {
	fmt.Printf("Details of the user is : %+v\n", p)
}
func Value(){
	fmt.Println("Hello")
}