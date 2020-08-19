package basic

func isAnagram(mapWord map[rune]int, word string) bool {
	for _, r := range word {
		v, ok := mapWord[r]
		if !ok || v <= 0 {
			return false
		}
		mapWord[r]--
	}

	for _, v := range mapWord {
		if v != 0 {
			return false
		}
	}

	return true
}

// RemoveAnagram removes the anagrams like ["code", "doce"] => ["code"]
func RemoveAnagram(texts []string) []string {
	for i := 0; i < len(texts); i++ {
		word := texts[i]

		mapWord := make(map[rune]int, len(word))
		for _, r := range word {
			mapWord[r]++
		}

		for j := i + 1; j < len(texts); j++ {
			if isAnagram(mapWord, texts[j]) {
				texts = append(texts[:j], texts[j+1:]...)
				i--
			}
		}

	}

	return texts
}
