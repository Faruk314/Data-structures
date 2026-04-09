package binarysearch

func BinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func SearchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

var target = 42

func guess(num int) int {
	if num == target {
		return 0
	} else if num > target {
		return -1
	} else {
		return 1
	}
}

func GuessNumber(n int) int {
	low := 1
	high := n
	for low <= high {
		mid := low + (high-low)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == -1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func MySqrt(x int) int {
	if x < 2 {
		return x
	}

	low := 1
	high := x
	ans := 0

	for low <= high {

		mid := low + (high-low)/2

		if mid <= x/mid {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return ans
}
