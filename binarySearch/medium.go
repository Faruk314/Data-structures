package binarysearch

import "slices"

func SeachMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	left := 0
	right := len(matrix) - 1

	for left <= right {
		mid := left + (right-left)/2
		first := matrix[mid][0]
		last := matrix[mid][len(matrix[mid])-1]

		if first == target || last == target {
			return true
		}

		if target > first && target < last {
			low := 0
			high := len(matrix[mid]) - 1

			for low <= high {
				innerMid := low + (high-low)/2

				if matrix[mid][innerMid] == target {
					return true
				} else if matrix[mid][innerMid] < target {
					low = innerMid + 1
				} else {
					high = innerMid - 1
				}
			}

			return false
		}

		if target > last {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return false
}

func CalcPiles(piles []int, h int) int {
	left := 1
	right := slices.Max(piles)
	res := right

	for left <= right {
		k := left + (right-left)/2

		hours := 0
		for _, pile := range piles {
			hours += (pile + k - 1) / k
		}

		if hours <= h {
			if k < res {
				res = k
			}
			right = k - 1
		} else {
			left = k + 1
		}
	}

	return res
}

func ShipWithingDays(weights []int, days int) int {
	minCap := 0
	maxCap := 0

	for _, weight := range weights {
		if weight > minCap {
			minCap = weight
		}
		maxCap += weight
	}

	for minCap <= maxCap {
		mid := minCap + (maxCap-minCap)/2

		foundDays := 1
		sum := 0

		for _, weight := range weights {
			if sum+weight > mid {
				foundDays++
				sum = 0
			}
			sum += weight
		}

		if foundDays > days {
			minCap = mid + 1
		} else {
			maxCap = mid - 1
		}
	}

	return minCap
}

func FindMin(nums []int) int {
	n := len(nums)

	left := 0
	right := n - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
