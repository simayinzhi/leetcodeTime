package main

func search(nums []int, target int) int {
	index := findMaxIndex(nums)
	if target >= nums[0] {
		return bSearch(nums[:index+1], target)
	} else {
		result := bSearch(nums[index+1:], target)
		if result != -1 {
			result += index + 1
		}
		return result
	}
}

func bSearch(nums []int, target int) int {
	if (len(nums)) == 0 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	mid := l + (r-l)/2
	for l <= r {
		mid = l + (r-l)/2
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func findMaxIndex(nums []int) int {
	if len(nums) < 2 {
		return len(nums) - 1
	}
	l := 0
	r := len(nums) - 1
	if nums[r] > nums[0] {
		return r
	}
	var mid int
	for r >= l {
		mid = l + (r-l)/2
		if nums[mid] < nums[0] {
			if nums[mid-1] >= nums[0] {
				return mid - 1
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return len(nums) - 1
}

func main() {
	println(search([]int{3, 1}, 1) == 1)
	println(search([]int{3, 1}, 3) == 0)
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0) == 4)
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3) == -1)
	println(search([]int{2}, 9) == -1)
	println(search([]int{4, 5, 1, 2, 3}, 6) == -1)
	println(search([]int{2, 3, 4}, 2) == 0)
	println(search([]int{3, 4, 5, 1, 2}, 2) == 4)
	println(search([]int{4, 5, 1, 2, 3}, 1) == 2)
	println(search([]int{3, 5, 1}, 1) == 2)
	println(search([]int{2, 3, 4}, 2) == 0)
	//println(findMaxIndex([]int{4, 5, 6, 7, 0, 1, 2}))
	//println(findMaxIndex([]int{3, 4, 5, 1, 2}) == 2)
	//println(findMaxIndex([]int{4, 5, 1, 2, 3}) == 1)
	//println(findMaxIndex([]int{3, 5, 1}) == 1)
	//println(findMaxIndex([]int{3, 4, 5, 1}) == 2)
	//println(findMaxIndex([]int{}))
	//println(findMaxIndex([]int{1}))
	//println(findMaxIndex([]int{1, 2, 3}))

}
