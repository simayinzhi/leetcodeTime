package main

import "fmt"

func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if flag != i {
				nums[flag] = nums[i]
			}
			flag++
		}
	}
	return flag
}

func main() {
	//nums := []int{1, 2, 3, 3, 3, 4, 5}
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k := removeElement(nums, 2)

	fmt.Println(k)
	fmt.Println(nums)
}
