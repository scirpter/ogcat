package gen

import "strings"

func GenUsernames(lengthOffset uint8, maxLen uint8, includeABC bool, includeSpecial bool, includeDigits bool) []string {
	var ret []string
	toGen := []string{}
	if includeABC {
		toGen = append(toGen, []string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
			"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
			"u", "v", "w", "x", "y", "z",
		}...)
	}
	if includeSpecial {
		toGen = append(toGen, []string{
			".", "_",
		}...)
	}
	if includeDigits {
		toGen = append(toGen, []string{
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		}...)
	}

	for length := lengthOffset; uint8(length) <= maxLen; length++ {
		for _, combination := range product(toGen, length) {
			username := strings.Join(combination, "")
			if !strings.Contains(username, "..") {
				ret = append(ret, username)
			}
		}
	}

	return ret
}

// helper function to generate combinations
func product(values []string, length uint8) [][]string {
	if length == 0 {
		return [][]string{{}}
	}

	var combinations [][]string

	for _, value := range values {
		for _, combo := range product(values, length-1) {
			combinations = append(combinations, append([]string{value}, combo...))
		}
	}

	return combinations
}
