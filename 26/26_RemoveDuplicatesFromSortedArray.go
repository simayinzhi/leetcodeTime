package main

import "fmt"

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) <= 1 {
		return len(nums)
	}
	flag := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		flag++
		nums[flag] = nums[i]
	}
	return flag + 1
}

func main() {
	//nums := []int{1, 2, 3, 3, 3, 4, 5}
	nums := []int{1}
	k := removeDuplicates(nums)

	fmt.Println(k)
	fmt.Println(nums)
}
