package main

import "fmt"

func main() {
	age := 10
	name := "Jack"
	rightHander := true

	fmt.Printf("%s is %d years old. right handed: %t", name, age, rightHander)

	ageInTenYears := age + 10

	fmt.Printf("in ten years %s will be %d years old", name, ageInTenYears)
	fmt.Println()

	isATenager := age >= 13
	fmt.Println(name, "is a teenager", isATenager)
}
