package alphametics

import (
	"errors"
	"math"
	"strings"
)

type letterInfo struct {
	val     int
	leading bool
}

type alphametic struct {
	letters    [26]letterInfo
	index      []byte
	mul        [26]int
	usedDigits [10]bool
}

func Solve(puzzle string) (map[string]int, error) {
	var a alphametic
	a.index = make([]byte, 0, 26)

	words := strings.FieldsFunc(puzzle, func(r rune) bool {
		return r == '+' || r == '=' || r == ' '
	})
	if len(words) < 2 {
		return nil, errors.New("invalid puzzle")
	}

	for i, w := range words {
		if len(w) == 0 {
			continue
		}
		first := w[0] - 'A'
		a.letters[first].leading = true
		sign := 1
		if i == len(words)-1 {
			sign = -1
		}

		l := len(w)
		for pos := 0; pos < l; pos++ {
			c := w[pos] - 'A'
			a.mul[c] += sign * int(math.Pow10(l-pos-1))
		}
	}

	for c := byte(0); c < 26; c++ {
		if a.mul[c] != 0 || a.letters[c].leading {
			a.index = append(a.index, c)
		}
	}

	if !a.solveRec(0, 0) {
		return nil, errors.New("no solution")
	}

	res := make(map[string]int, len(a.index))
	for _, c := range a.index {
		res[string('A'+c)] = a.letters[c].val
	}
	return res, nil
}

func (a *alphametic) solveRec(i int, sum int) bool {
	if i == len(a.index) {
		return sum == 0
	}

	c := a.index[i]
	info := &a.letters[c]

	startDigit := 0
	if info.leading {
		startDigit = 1
	}

	for d := startDigit; d <= 9; d++ {
		if a.usedDigits[d] {
			continue
		}

		a.usedDigits[d] = true
		info.val = d
		if a.solveRec(i+1, sum+d*a.mul[c]) {
			return true
		}

		a.usedDigits[d] = false
	}

	return false
}
