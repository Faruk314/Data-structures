package binarysearch

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
