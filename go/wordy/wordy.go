package wordy

import (
	"strconv"
	"strings"
)

var operation = map[string]func(int, int) int{
	"plus":          func(x, y int) int { return x + y },
	"minus":         func(x, y int) int { return x - y },
	"multiplied_by": func(x, y int) int { return x * y },
	"divided_by":    func(x, y int) int { return x / y },
}

func Answer(question string) (int, bool) {
	tokens := parseQuestion(question)
	if tokens == nil {
		return 0, false
	}

	x, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, false
	}

	for i := 1; i < len(tokens); i += 2 {
		op, ok := operation[tokens[i]]
		val, err := strconv.Atoi(tokens[i+1])
		if !ok || err != nil {
			return 0, false
		}

		x = op(x, val)
	}

	return x, true
}

func parseQuestion(question string) []string {
	question = strings.TrimPrefix(question, "What is ")
	question = strings.TrimSuffix(question, "?")
	question = strings.ReplaceAll(question, "multiplied by", "multiplied_by")
	question = strings.ReplaceAll(question, "divided by", "divided_by")

	tokens := strings.Fields(question)
	if len(tokens) == 0 || len(tokens)%2 == 0 {
		return nil
	}

	return tokens
}
