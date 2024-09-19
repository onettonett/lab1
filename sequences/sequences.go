package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) []int {
	newSlice := []int{}
	for _, number := range slice {
		newSlice = append(newSlice, f(number))
	}
	return newSlice
}

func mapArray(f func(a int) int, array [5]int) [5]int {
	var newArray [5]int
	for i, number := range array {
		newArray[i] = f(number)
	}
	return newArray
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int(mapSlice(addOne, slice1))
	fmt.Println(slice2)
	intsArray := [5]int{1, 2, 3, 4, 5}
	intsArray2 := [5]int(mapArray(addOne, intsArray))
	fmt.Println(intsArray2)
	slice3 := double(slice1)
	fmt.Println(slice3)
}
