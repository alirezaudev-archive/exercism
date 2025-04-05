package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	result := make([]rune, len(plain))
	for i, r := range plain {
		switch {
		case unicode.IsLower(r):
			result[i] = 'a' + (r-'a'+rune(shiftKey))%26
		case unicode.IsUpper(r):
			result[i] = 'A' + (r-'A'+rune(shiftKey))%26
		default:
			result[i] = r
		}
	}
	return string(result)
}
