package arrayshashing

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		chars := strings.Split(s, "")

		sort.Strings(chars)

		key := strings.Join(chars, "")

		groups[key] = append(groups[key], s)

	}

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
