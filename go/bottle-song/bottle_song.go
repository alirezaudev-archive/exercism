package bottlesong

import (
	"fmt"
	"strings"
)

var nums = []string{
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
}

func Recite(startBottles, takeDown int) []string {
	ret := make([]string, 0)
	count := 0

	tmp := "%s hanging on the wall,"
	for i := startBottles; i > 0 && count < takeDown; i -= 1 {
		count += 1
		bottleTxt := greenBottle(i - 1)
		ret = append(ret, fmt.Sprintf(tmp, bottleTxt))
		ret = append(ret, fmt.Sprintf(tmp, bottleTxt))
		ret = append(ret, "And if one green bottle should accidentally fall,")
		if i == 1 {
			ret = append(ret, "There'll be no green bottles hanging on the wall.")
		} else {
			ret = append(ret, fmt.Sprintf("There'll be %s hanging on the wall.", strings.ToLower(greenBottle(i-2))))
		}
		if count > 0 && count < takeDown {
			ret = append(ret, "")
		}
	}
	return ret
}

func greenBottle(i int) string {
	txt := nums[i] + " green bottle"
	if i > 0 {
		txt += "s"
	}

	return txt
}
