package main

import "fmt"

func basicCalculation(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums)-1; j++ {

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}

}
func twoSum(nums []int, target int) [][]int {

	traversed := make(map[int]int)
	returnSlices := make([][]int, 0)
	for index, value := range nums {
		traversed[value] = index
		if val, ok := traversed[target-value]; ok {
			returnSlices = append(returnSlices, []int{val, index})
		}

	}

	return returnSlices
}

func main() {

	// Let there be a slice of integeres as

	s := []int{20, 11, 7, 6, 15, 3, 2}
	// and
	target := 15

	// Send the parameters to twoSun function to get the two indexes of the Slice
	returnedSlice := twoSum(s, target)
	fmt.Println(returnedSlice)
	for _, v := range returnedSlice {
		var sumOfIndexValues = 0
		for _, value := range v {

			sumOfIndexValues += s[value]

		}
		if sumOfIndexValues == target {
			
			fmt.Println("Our logic is working fine for Indexes: ",v)
		} else {
			fmt.Println("Evaluate your logic for Indexes: ",v)
		}
	}

}
