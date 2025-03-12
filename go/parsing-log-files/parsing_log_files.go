package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)].*`).MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<([~=*\-])*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if regexp.MustCompile(`".*((?i)password).*"`).MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d*`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	for i, line := range lines {
		re := regexp.MustCompile(`User\s+(\w+)\s`).FindStringSubmatch(line)
		if len(re) > 1 {
			lines[i] = fmt.Sprintf("[USR] %s %s", re[1], line)
		}
	}

	return lines
}
