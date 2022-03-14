package main

import "fmt"

func searchRange(nums []int, target int) []int {
	index := binarySearch(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}
	right := index
	for ; right < len(nums)-1; right++ {
		if nums[right] != nums[right+1] {
			break
		}
	}
	left := index
	for ; left > 0; left-- {
		if nums[left] != nums[left-1] {
			break
		}
	}

	return []int{left, right}
}

func binarySearch(numbers []int, target int) int {
	if len(numbers) == 0 {
		return -1
	}
	left := 0
	right := len(numbers) - 1
	middle := left + (right-left)/2
	for left <= right {
		if target == numbers[middle] {
			return middle
		} else if target > numbers[middle] {
			left = middle + 1
			middle = left + (right-left)/2
		} else {
			right = middle - 1
			middle = left + (right-left)/2
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 2, 3, 4, 4}, 4))
	fmt.Println(binarySearch([]int{1, 2, 2, 3, 4, 4}, 2))
	fmt.Println(binarySearch([]int{4, 4}, 2))
	fmt.Println(searchRange([]int{1, 2, 2, 3, 4, 4}, 4))
	fmt.Println(searchRange([]int{1, 2, 2, 3, 4, 4}, 2))
	fmt.Println(searchRange([]int{1, 2, 2, 3, 4, 4}, 3))
	fmt.Println(searchRange([]int{4, 4}, 4))
	fmt.Println(searchRange([]int{4, 4}, 1))
	fmt.Println(searchRange([]int{}, 1))
	fmt.Println(searchRange([]int{1}, 1))
	fmt.Println(searchRange([]int{1}, 7))

}
