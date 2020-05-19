package main

import "fmt"

func main() {
	slicer := []string{"f", "u", "j", "i", "y", "a"}
	slicers := []string{"m", "i", "y", "a", "g", "i"}
	slicers = append(slicers, slicer...)

	fmt.Println(slicers)
}
