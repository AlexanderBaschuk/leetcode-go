package main

import (
	"fmt"
	twoSumPkg "leetcode/1_twoSum"
)

func main() {
	fmt.Println("Hello, World!")
	result := twoSumPkg.TwoSum([]int{1, 2, 3, 5}, 5)
	fmt.Println(result[0], result[1])
}
