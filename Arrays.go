package main

import "fmt"

func main() {
	//var values[5] int = [5] int {5, 25, 125, 625, 3125}
	var values = [5]int{5, 25, 125, 625, 3125}

	names := [4]string{"Joker", "Futaba", "Morgana", "Noir"}
	names[3] = "Violet"

	fmt.Println(values, len(values))
	fmt.Println(names, len(names))

	// Slices (Use arrays under the hood)
	var scores = []int{0000, 0001, 0010, 0100}
	scores[1] = 150
	scores = append(scores, 1010)

	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeA := names[1:4] // Doesn't include poss 4 element
	rangeB := names[2:]  // Includes the last element
	rangeC := names[:3]

	fmt.Println(rangeA, rangeB, rangeC)
	fmt.Printf("The type of rangeA is : %T \n", rangeA)
	rangeA = append(rangeA, "Skull")
	fmt.Println(rangeA)

}
