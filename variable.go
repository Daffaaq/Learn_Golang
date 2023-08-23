package main

import "fmt"
func main() {
	var name string

	name = "Daffa Aqila Rahmatullah"
	fmt.Println(name)

	name = "Daffa Aqila"
	fmt.Println(name)

	country := "Indonesia"
	fmt.Println(country,name)

	var (
		firstName = "Daffa Aqila"
		lastName = "Rahmatullah"
	)
	fmt.Println(firstName+"",lastName) //Daffa Aqila Rahmatullah
	fmt.Println(firstName+lastName) //Daffa AqilaRahmatullah
	fmt.Println(firstName,lastName) //Daffa Aqila Rahmatullah
}