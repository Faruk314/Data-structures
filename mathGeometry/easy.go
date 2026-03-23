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
