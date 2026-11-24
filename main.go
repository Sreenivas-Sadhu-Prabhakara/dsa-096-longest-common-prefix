package main

import (
	"fmt"
	"strings"
)

// LongestCommonPrefix returns the longest prefix shared by every string.
// Start from the first string and shrink it against each other string.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs[1:] {
		for !strings.HasPrefix(s, prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

func main() {
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
}
