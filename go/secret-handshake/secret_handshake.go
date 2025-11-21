package secret

var flags = []uint{
	0b00001,
	0b00010,
	0b00100,
	0b01000,
}

var actions = map[uint]string{
	0b00001: "wink",
	0b00010: "double blink",
	0b00100: "close your eyes",
	0b01000: "jump",
}

func Handshake(code uint) []string {
	var result []string

	for _, flag := range flags {
		if code&flag != 0 {
			result = append(result, actions[flag])
		}
	}

	if code&0b10000 != 0 {
		reverse(result)
	}

	return result
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
