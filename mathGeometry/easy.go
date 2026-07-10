package mathgeometry

// first solution
func RomanToInt(s string) int {
	romanToIntMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	sum := 0
	left := 0
	right := 1

	for {
		if left+1 > len(s) {
			break
		}

		roman1 := s[left]
		if left+1 == len(s) {

			if num, ok := romanToIntMap[string(roman1)]; ok {
				sum += num
			}
			break
		}

		roman2 := s[right]
		res := string([]byte{roman1, roman2})
		if num, ok := romanToIntMap[res]; ok {
			sum += num

			left += 2
			right += 2
		} else {

			if num, ok := romanToIntMap[string(roman1)]; ok {
				sum += num
			}

			left++
			right++

		}

	}

	return sum
}

func romanToInt(s string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && values[s[i]] < values[s[i+1]] {
			total -= values[s[i]]
		} else {
			total += values[s[i]]
		}
	}

	return total
}

func ConvertToTitle(columnNum int32) string {
	result := ""

	for columnNum > 0 {
		columnNum--

		char := byte(columnNum%26 + 'A')

		result = string(char) + result

		columnNum = columnNum / 26
	}

	return result
}

func GcdOfStrings(str1 string, str2 string) string {
	match := ""

	shorter := str1
	bigger := str2

	if len(shorter) > len(bigger) {
		shorter, bigger = bigger, shorter
	}

	for i := len(shorter); i > 0; i-- {
		key := shorter[:i]
		didMatch := false

		if len(bigger)%len(key) != 0 || len(shorter)%len(key) != 0 {
			continue
		}

		for j := 0; j < len(bigger); j += len(key) {
			if bigger[j:len(key)+j] == key {
				didMatch = true
			} else {
				didMatch = false
				break
			}
		}

		if didMatch {
			for j := 0; j < len(shorter); j += len(key) {
				if shorter[j:len(key)+j] == key {
					didMatch = true
				} else {
					didMatch = false
					break
				}
			}
		}

		if didMatch {
			match = key
			break
		}
	}

	return match
}

func diagonalSum(mat [][]int) int {
	width, height := len(mat[0]), len(mat)
	dir := [2][2]int{{0, 0}, {0, width - 1}}

	totalSum := 0

	for idx, d := range dir {

		row := d[0]
		col := d[1]

		for row >= 0 && row < height && col >= 0 && col < width {
			totalSum += mat[row][col]

			row++

			if idx == 0 {
				col = row
			} else {
				col = width - row - 1
			}
		}

	}

	if width%2 == 0 {
		return totalSum
	}

	return totalSum - mat[height/2][width/2]
}

func largestOddNumber(num string) string {
	result := ""

	for i := len(num) - 1; i >= 0; i-- {
		if num[i]%2 != 0 {
			result = num[:i+1]
			break
		}
	}

	return result
}

func transpose(matrix [][]int) [][]int {
	height, width := len(matrix), len(matrix[0])

	res := make([][]int, width)

	for i := 0; i < len(res); i++ {
		res[i] = make([]int, height)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			res[j][i] = matrix[i][j]
		}
	}

	return res
}
