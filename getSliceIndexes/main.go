package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) <2{
		return []int {}
	}

	for i:=0; i< len(nums); i++{
		for j:=i+1; j<len(nums)-1; j++{

			if nums[i]+nums[j] == target{
				return []int{i, j}
			}
		}
	}
	return []int {}
    
}

func main(){

	// Let there be a slice of integeres as 

	s:=[]int{2,7,11,15}
	// and 
	target:=9

	// Send the parameters to twoSun function to get the two indexes of the Slice 
	returnedSlice:= twoSum(s, target)


	var sumOfIndexValues=0
	for _,v:=range returnedSlice{
		sumOfIndexValues +=s[v]
	}
	if sumOfIndexValues==target{
		fmt.Println(returnedSlice)
		fmt.Println("Our logic is working fine!")
	}else{
		fmt.Println(returnedSlice)
		fmt.Println("Evaluate your logic!")
	}
}

