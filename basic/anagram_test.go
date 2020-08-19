package basic

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	tests := []struct {
		Input  []string
		Output []string
	}{
		{
			Input:  []string{"code", "doce", "cedo", "anagram", "managra", "angrams", "odec"},
			Output: []string{"code", "anagram", "angrams"},
		},
		// add any other testcase
	}

	for _, test := range tests {
		result := RemoveAnagram(test.Input)
		if !isEqual(result, test.Output) {
			t.Fatal(fmt.Sprintf("Anagram should be: %+v instread of %+v", test.Output, result))
		}
	}
}

func isEqual(str, str1 []string) bool {
	if len(str) != len(str1) {
		return false
	}

	for i := 0; i < len(str); i++ {
		if str[i] != str1[i] {
			return false
		}
	}

	return true
}
