package main

import "fmt"

/*
	You are delivering packages to customers. You have 100 packages to deliver.

	1. Print the number of package you are going to deliver For example: "I have 100 packages to deliver"

	2. You have delivered 20 packages. Print the remaining packages to deliver

	3. The packages are going to be distributed equally between 4 customers. Print how many packages each customer receives (while mentioning the number of customers)
*/

func main() {
	var customers = 4
	var packagesToDeliver = 100

	fmt.Println("I have", packagesToDeliver, "packages to deliver")

	var deliveredPackages = 20
	packagesToDeliver -= deliveredPackages
	fmt.Printf("I have delivered %v packages\n", deliveredPackages)
	fmt.Printf("Remaining packages to deliver: %v\n", packagesToDeliver)

	packagesPerCustomer := packagesToDeliver / customers
	fmt.Printf("Packages are going to be distributed equally between %v customers. This means %v packages per customer", customers, packagesPerCustomer)
}
