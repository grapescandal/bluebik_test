package main

import (
	"bluebik_test/answer1"
	"bluebik_test/answer2"
	"bluebik_test/answer3"
	"fmt"
)

func main() {

	//Question1
	fmt.Println("----------------------Question1----------------------")
	arr := []int{1, 2, 4, 5, 0, 6, 3, 8, 9, 10}

	output := answer1.FindMissingElement(arr)
	fmt.Println("Output:", output)

	//Question2
	fmt.Println("----------------------Question2----------------------")

	str := "1234567890"

	containsOnlyDigits := answer2.ContainsOnlyDigits(str)

	if containsOnlyDigits {
		fmt.Println("String contains only digits")
	} else {
		fmt.Println("String does not contain only digits")
	}

	//Question3
	fmt.Println("----------------------Question3----------------------")

	numRows := 6
	pascalTriangle := answer3.GeneratePascalTriangle(numRows)

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(pascalTriangle[i][j], " ")
		}
		fmt.Println()
	}
}
