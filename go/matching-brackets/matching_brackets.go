package brackets

func Bracket(input string) bool {
	opens := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	closes := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]rune, len(input))
	i := 0
	for _, char := range input {
		if _, isOpen := opens[char]; isOpen {
			stack[i] = char
			i++
			continue
		}
		if _, isClose := closes[char]; !isClose {
			continue
		}

		if i == 0 || stack[i-1] != closes[char] {
			return false
		}

		stack[i-1] = 0
		i--
	}

	return i == 0
}
