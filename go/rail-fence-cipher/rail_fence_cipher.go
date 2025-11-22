package railfence

import (
	"strings"
)

func Encode(msg string, rails int) string {
	if rails <= 1 || len(msg) == 0 {
		return msg
	}

	railStrings := make([]strings.Builder, rails)
	cycle := 2 * (rails - 1)
	for i := 0; i < len(msg); i++ {
		t := i % cycle

		var railIndex int
		if t < rails {
			railIndex = t
		} else {
			railIndex = cycle - t
		}

		railStrings[railIndex].WriteByte(msg[i])
	}

	var result strings.Builder
	for _, rail := range railStrings {
		result.WriteString(rail.String())
	}

	return result.String()
}

func Decode(message string, rails int) string {
	if rails <= 1 || len(message) == 0 {
		return message
	}

	n := len(message)
	cycle := 2 * (rails - 1)

	counts := make([]int, rails)
	for i := 0; i < n; i++ {
		t := i % cycle
		var railIndex int
		if t < rails {
			railIndex = t
		} else {
			railIndex = cycle - t
		}
		counts[railIndex]++
	}

	railStrings := make([]strings.Builder, rails)
	pos := 0
	for r := 0; r < rails; r++ {
		if counts[r] > 0 {
			railStrings[r].Grow(counts[r])
			railStrings[r].WriteString(message[pos : pos+counts[r]])
			pos += counts[r]
		}
	}

	railReadPos := make([]int, rails)
	var result strings.Builder
	result.Grow(n)

	for i := 0; i < n; i++ {
		t := i % cycle
		var railIndex int
		if t < rails {
			railIndex = t
		} else {
			railIndex = cycle - t
		}

		b := railStrings[railIndex].String()[railReadPos[railIndex]]
		railReadPos[railIndex]++
		result.WriteByte(b)
	}

	return result.String()
}
