package cipher

import (
	"math/rand"
	"strings"
	"time"
)

type shiftCipher struct {
	shift int
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(shift int) Cipher {
	if shift == 0 || shift < -25 || shift > 25 {
		return nil
	}
	return &shiftCipher{shift: shift}
}

func (c *shiftCipher) Encode(input string) string {
	return shiftString(input, c.shift)
}

func (c *shiftCipher) Decode(input string) string {
	return shiftString(input, -c.shift)
}

type vigenereCipher struct {
	key string
}

func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}

	if len(key) == 0 {
		return &vigenereCipher{key: randomKey(100)}
	}
	for _, ch := range key {
		if ch < 'a' || ch > 'z' {
			return nil
		}
	}
	if strings.Trim(key, "a") == "" {
		return nil
	}
	return &vigenereCipher{key: key}
}

func (v *vigenereCipher) Encode(input string) string {
	input = normalize(input)
	res := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = shiftChar(input[i], int(v.key[i%len(v.key)]-'a'))
	}
	return string(res)
}

func (v *vigenereCipher) Decode(input string) string {
	input = normalize(input)
	res := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = shiftChar(input[i], -int(v.key[i%len(v.key)]-'a'))
	}
	return string(res)
}

func normalize(s string) string {
	s = strings.ToLower(s)
	res := strings.Builder{}
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			res.WriteRune(ch)
		}
	}
	return res.String()
}

func shiftString(s string, shift int) string {
	s = normalize(s)
	res := make([]byte, len(s))
	for i := range s {
		res[i] = shiftChar(s[i], shift)
	}
	return string(res)
}

func shiftChar(c byte, shift int) byte {
	return byte((int(c-'a')+shift+26)%26 + 'a')
}

func randomKey(n int) string {
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = byte('a' + rand.Intn(26))
	}
	return string(res)
}
