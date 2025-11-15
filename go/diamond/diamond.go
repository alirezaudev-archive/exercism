package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("invalid")
	}

	if char == 'A' {
		return "A", nil
	}

	var b strings.Builder
	n := int(char - 'A')

	for i := 0; i <= n; i++ {
		b.WriteString(strings.Repeat(" ", n-i))
		b.WriteByte(byte('A' + i))
		if i > 0 {
			b.WriteString(strings.Repeat(" ", i*2-1))
			b.WriteByte(byte('A' + i))
		}
		b.WriteString(strings.Repeat(" ", n-i))
		b.WriteByte('\n')
	}

	for i := n - 1; i >= 0; i-- {
		b.WriteString(strings.Repeat(" ", n-i))
		b.WriteByte(byte('A' + i))
		if i > 0 {
			b.WriteString(strings.Repeat(" ", i*2-1))
			b.WriteByte(byte('A' + i))
		}
		b.WriteString(strings.Repeat(" ", n-i))
		if i != 0 {
			b.WriteByte('\n')
		}
	}

	return b.String(), nil
}
