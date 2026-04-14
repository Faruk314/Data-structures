package twopointers

import (
	"slices"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-3; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := len(nums) - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--

				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

func MaxArea(heights []int) int {
	leftPointer := 0
	rightPointer := len(heights) - 1
	biggestSurface := 0

	for leftPointer < rightPointer {

		leftHeight := heights[leftPointer]

		rightHeight := heights[rightPointer]

		height := min(leftHeight, rightHeight)

		width := rightPointer - leftPointer

		surface := height * width

		if surface > biggestSurface {
			biggestSurface = surface
		}

		if rightHeight >= leftHeight {
			leftPointer++
		} else {
			rightPointer--
		}

	}

	return biggestSurface
}

func BoatsToSavePeople(people []int, limit int) int {
	slices.Sort(people)
	count := 0
	left := 0
	right := len(people) - 1

	for left <= right {

		if people[left]+people[right] <= limit {
			left++
		}

		right--
		count++
	}

	return count
}

func LongestMountain(arr []int) int {
	res := 0

	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			left, right := i, i

			for left > 0 && arr[left] > arr[left-1] {
				left--
			}

			for right+1 < len(arr) && arr[right] > arr[right+1] {
				right++
			}

			total := right - left + 1

			if total > res {
				res = total
			}
		}
	}

	return res
}
