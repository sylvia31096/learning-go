package main

import (
	"fmt"
	"strconv"
)

func main() {

	// prompt user enter the set of values to sort
	fmt.Println("Enter list of values for sorting or stop to stop entering")

	var sortelement string //create variables to hold the entered value
	var sortelementint int //create variable to hold the entered value as an int
	sortlist := make([]int, 0, 3)

	//allow user to enter value until the list reaches 10 or the user stops
	for ("stop" != sortelement) && (len(sortlist) < 10) {

		fmt.Scan(&sortelement)
		sortelementint, _ = strconv.Atoi(sortelement) //convert entered value to int
		sortlist = append(sortlist, sortelementint)   //append to slice
	}

	fmt.Println("sorted list:\n", BubbleSort(sortlist)) //print sorted list

}

func BubbleSort(sortlist []int) []int {

	/* sorts a list of numbers using
	bubble sort algorithm */

	maxlen := len(sortlist)

	for z := 0; z < maxlen; z++ {

		for k := 0; k < maxlen-1; k++ {
			Swap(sortlist, k)
		}
	}

	return sortlist

}

func Swap(sortlist []int, i int) {

	/* swaps the order of two numbers if
	the previous number is greater than the
	next number */

	j := i + 1
	var temp int
	if sortlist[i] > sortlist[j] {
		temp = sortlist[i]
		sortlist[i] = sortlist[j]
		sortlist[j] = temp
	}
}
