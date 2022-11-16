package main

import (
	"fmt"
	"spe-skill-test/challenges"
)

func main() {
	// Narcissistic Number
	isNarcissistic := challenges.Narcissistic(153)
	fmt.Println(isNarcissistic)
	isNarcissistic = challenges.Narcissistic(111)
	fmt.Println(isNarcissistic)

	// ======================================================
	fmt.Println()
	// ======================================================

	// Parity Outlier
	arr := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	fmt.Println(arr)
	fmt.Println(challenges.ParityOutlier(arr))
	arr = []int{160, 3, 1719, 19, 11, 13, -21}
	fmt.Println(arr)
	fmt.Println(challenges.ParityOutlier(arr))
	arr = []int{11, 13, 15, 19, 9, 13, -21}
	fmt.Println(arr)
	fmt.Println(challenges.ParityOutlier(arr))

	// ======================================================
	fmt.Println()
	// ======================================================

	// Needle In The Haystack
	fmt.Println(challenges.FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue"))

	// ======================================================
	fmt.Println()
	// ======================================================

	// The Blue Ocean Reverse
	res := challenges.BlueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1})
	fmt.Println(res)
	res = challenges.BlueOcean([]int{1, 5, 5, 5, 5, 3}, []int{5})
	fmt.Println(res)
}
