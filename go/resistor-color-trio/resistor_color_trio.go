package resistorcolortrio

import (
	"strconv"
	"strings"
)

var bands = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

var labels = map[int]string{
	0: " ohms",
	3: " kiloohms",
	6: " megaohms",
	9: " gigaohms",
}

func Label(colors []string) string {
	value := bands[colors[0]]*10 + bands[colors[1]]
	zerosCount := bands[colors[2]]

	if value > 10 && value%10 == 0 {
		zerosCount++
		value /= 10
	}

	labelKey := 0
	switch {
	case zerosCount >= 9:
		labelKey = 9
	case zerosCount >= 6:
		labelKey = 6
	case zerosCount >= 3:
		labelKey = 3
	}

	return strconv.Itoa(value) + strings.Repeat("0", zerosCount-labelKey) + labels[labelKey]
}
