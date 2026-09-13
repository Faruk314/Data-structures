package mathgeometry

import (
	"fmt"
	"math"
	"strings"
)

func IntToRoman(num int) string {
	type romanMap struct {
		val int
		sym string
	}

	romanSymbols := []romanMap{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder

	for _, entry := range romanSymbols {
		for num >= entry.val {
			num -= entry.val
			result.WriteString(entry.sym)
		}
	}

	return result.String()
}

func checkPowersOfThree(n int) bool {
	var currNum float64
	nums := []float64{}
	idx := 0

	for {
		currNum = math.Pow(3, float64(idx))

		if currNum > float64(n) {
			break
		}

		nums = append(nums, currNum)
		idx++
	}

	totalSum := nums[len(nums)-1]

	if totalSum == float64(n) {
		return true
	}

	for i := len(nums) - 2; i >= 0; i-- {

		if totalSum+nums[i] > float64(n) {
			continue
		}

		totalSum += nums[i]

		if totalSum == float64(n) {
			return true
		}
	}

	return false
}

func isRobotBounded(instructions string) bool {
	x := 0
	y := 0
	dir := 'N'

	for _, instruction := range instructions {
		if instruction == 'G' {
			switch dir {
			case 'N':
				y += 1
			case 'S':
				y -= 1
			case 'W':
				x -= 1
			case 'E':
				x += 1
			}
		} else if instruction == 'L' {
			switch dir {
			case 'N':
				dir = 'W'
			case 'W':
				dir = 'S'
			case 'S':
				dir = 'E'
			case 'E':
				dir = 'N'
			}
		} else if instruction == 'R' {
			switch dir {
			case 'N':
				dir = 'E'
			case 'E':
				dir = 'S'
			case 'S':
				dir = 'W'
			case 'W':
				dir = 'N'
			}
		}
	}

	return (x == 0 && y == 0) || dir != 'N'
}

func robotSim(commands []int, obstacles [][]int) int {
	maxDistance := 0
	x := 0
	y := 0
	dir := 'N'
	obstaclesMap := make(map[string]struct{})

	for _, obstacle := range obstacles {
		key := fmt.Sprintf("%d,%d", obstacle[0], obstacle[1])
		obstaclesMap[key] = struct{}{}
	}

	for _, num := range commands {
		if num == -2 {
			switch dir {
			case 'N':
				dir = 'W'
			case 'W':
				dir = 'S'
			case 'S':
				dir = 'E'
			case 'E':
				dir = 'N'
			}
		} else if num == -1 {
			switch dir {
			case 'N':
				dir = 'E'
			case 'E':
				dir = 'S'
			case 'S':
				dir = 'W'
			case 'W':
				dir = 'N'
			}
		} else {
			for i := 0; i < num; i++ {
				currX := x
				currY := y

				switch dir {
				case 'N':
					currY += 1
				case 'S':
					currY -= 1
				case 'W':
					currX -= 1
				case 'E':
					currX += 1
				}

				key := fmt.Sprintf("%d,%d", currX, currY)

				if _, exists := obstaclesMap[key]; exists {
					break
				}

				x = currX
				y = currY
				maxDistance = max(maxDistance, int(math.Pow(float64(x), 2)+math.Pow(float64(y), 2)))
			}
		}
	}

	return maxDistance
}
