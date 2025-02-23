package main

import "fmt"

func main() {
	age := 44          // regular variable
	agePointer := &age // pointer

	fmt.Println("Age: ", age)
	fmt.Println("Age pointer: ", agePointer)
	fmt.Println("Age pointer value: ", *agePointer)

	adultYears := getAdultYears(age)
	fmt.Println("Adult age: ", adultYears)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult age pointer: ", age)
}

func getAdultYears(age int) int {
	return age - 18
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
