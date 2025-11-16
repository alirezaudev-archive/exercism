package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	multi := len(files) > 1
	var results []string

	for _, file := range files {
		results = append(results, searchFile(file, pattern, flags, multi)...)
	}

	return results
}

func searchFile(filename, pattern string, flags []string, multi bool) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var results []string
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		if !matches(line, pattern, flags) {
			continue
		}

		if hasFlag(flags, "-l") {
			return []string{filename}
		}

		output := line
		if hasFlag(flags, "-n") {
			output = fmt.Sprintf("%d:%s", lineNum, output)
		}
		if multi {
			output = filename + ":" + output
		}

		results = append(results, output)
	}

	return results
}

func matches(line, pattern string, flags []string) bool {
	if hasFlag(flags, "-i") {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	var matched bool
	if hasFlag(flags, "-x") {
		matched = line == pattern
	} else {
		matched = strings.Contains(line, pattern)
	}

	if hasFlag(flags, "-v") {
		matched = !matched
	}

	return matched
}

func hasFlag(flags []string, flag string) bool {
	for _, f := range flags {
		if f == flag {
			return true
		}
	}
	return false
}
