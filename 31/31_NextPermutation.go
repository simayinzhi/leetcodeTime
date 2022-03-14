package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if i > 0 && nums[i-1] < nums[i] {
			nextSeq(nums[i-1:])
			return
		}
		if i == 0 {
			sort.Ints(nums)
			return
		}
	}
}

func nextSeq(nums []int) {
	//swap(nums, 0, len(nums)-1)
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[0] {
			swap(nums, 0, i)
			break
		}
	}
	// find next number bigger than 0
	sort.Ints(nums[1:len(nums)])
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverseCount(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
	}
	return count
}
func maxReverseCount(size int) int {
	size = size - 1
	return (size * (1 + size)) / 2.0
}

func main() {
	//nums := []int{5, 1, 4, 3, 3}
	//nums := []int{1, 2, 3}
	//nums := []int{3, 2, 1}
	//nums := []int{2, 3, 1}
	nums := []int{0}
	//nums := []int{1, 1, 55555}
	nextPermutation(nums)
	fmt.Println(nums)
}
