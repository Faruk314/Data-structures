package mathgeometry

import "strings"

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
