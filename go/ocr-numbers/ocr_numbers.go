package ocr

import "strings"

var table = map[string]byte{
	" _ | ||_|": '0',
	"     |  |": '1',
	" _  _||_ ": '2',
	" _  _| _|": '3',
	"   |_|  |": '4',
	" _ |_  _|": '5',
	" _ |_ |_|": '6',
	" _   |  |": '7',
	" _ |_||_|": '8',
	" _ |_| _|": '9',
}

func Recognize(s string) []string {
	var result []string

	lines := strings.Split(s, "\n")
	for i := 1; i < len(lines); i += 4 {
		result = append(result, recognizeDigit(lines[i:i+3]))
	}
	return result
}

func recognizeDigit(line []string) string {
	var tmp strings.Builder
	for j := 0; j < len(line[0]); j += 3 {
		tmp.WriteByte(getDigit(line[0][j:j+3] + line[1][j:j+3] + line[2][j:j+3]))
	}

	return tmp.String()
}

func getDigit(s string) byte {
	if d, ok := table[s]; ok {
		return d
	}
	return '?'
}
