package say

var words = []string{
	"", "one", "two", "three", "four", "five", "six",
	"seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen",
	"sixteen", "seventeen", "eighteen", "nineteen",
}

var tens = []string{
	"", "", "twenty", "thirty", "forty", "fifty",
	"sixty", "seventy", "eighty", "ninety",
}

var scales = map[int64]string{
	1000000000: "billion",
	1000000:    "million",
	1000:       "thousand",
	100:        "hundred",
}

var scaleOrder = []int64{1000000000, 1000000, 1000, 100}

func Say(n int64) (string, bool) {
	if n < 0 || n > 999999999999 {
		return "", false
	}

	if n == 0 {
		return "zero", true
	}

	return convertNumber(n), true
}

func convertNumber(n int64) string {
	if n == 0 {
		return ""
	}

	if n < 20 {
		return words[n]
	}

	if n < 100 {
		ten := n / 10
		remainder := n % 10
		if remainder == 0 {
			return tens[ten]
		}
		return tens[ten] + "-" + words[remainder]
	}

	for _, scale := range scaleOrder {
		if n >= scale {
			divisor := n / scale
			remainder := n % scale
			result := convertNumber(divisor) + " " + scales[scale]
			if remainder == 0 {
				return result
			}
			return result + " " + convertNumber(remainder)
		}
	}

	return ""
}
