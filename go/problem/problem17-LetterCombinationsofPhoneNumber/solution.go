func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// Mapping of digits to letters
	digitToLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var combinations []string
	var currentCombination []byte

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(digits) {
			combinations = append(combinations, string(currentCombination))
			return
		}
		letters := digitToLetters[digits[index]]
		for i := 0; i < len(letters); i++ {
			currentCombination = append(currentCombination, letters[i])
			backtrack(index + 1)
			currentCombination = currentCombination[:len(currentCombination)-1] // backtrack
		}
	}
	backtrack(0)
	return combinations
}
