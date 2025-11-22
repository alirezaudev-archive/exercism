package markdown

import (
	"fmt"
	"strings"
)

func Render(markdown string) string {
	markdown = applyInlineTags(markdown)
	markdown = replaceHeaders(markdown)
	markdown = replaceLists(markdown)
	markdown = replaceParagraphs(markdown)
	return markdown
}

func applyInlineTags(md string) string {
	md = strings.Replace(md, "__", "<strong>", 1)
	md = strings.Replace(md, "__", "</strong>", 1)
	md = strings.Replace(md, "_", "<em>", 1)
	md = strings.Replace(md, "_", "</em>", 1)
	return md
}

func replaceHeaders(md string) string {
	lines := strings.Split(md, "\n")
	var result strings.Builder

	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			result.WriteString(line)
			result.WriteString("\n")
			continue
		}

		level := 0
		for level < len(line) && line[level] == '#' {
			level++
		}

		if level > 6 || level >= len(line) || line[level] != ' ' {
			result.WriteString(line)
			result.WriteString("\n")
			continue
		}

		content := strings.TrimSpace(line[level:])
		result.WriteString(fmt.Sprintf("<h%d>%s</h%d>", level, content, level))
		result.WriteString("\n")
	}

	return result.String()
}

func replaceLists(md string) string {
	lines := strings.Split(md, "\n")
	var result strings.Builder
	inList := false

	for _, line := range lines {
		if strings.HasPrefix(line, "* ") {
			if !inList {
				result.WriteString("<ul>")
				inList = true
			}
			content := strings.TrimSpace(line[2:])
			result.WriteString("<li>" + content + "</li>")
		} else {
			if inList {
				result.WriteString("</ul>")
				result.WriteString("\n")
				inList = false
			}
			result.WriteString(line)
			result.WriteString("\n")
		}
	}

	if inList {
		result.WriteString("</ul>")
		result.WriteString("\n")
	}

	return result.String()
}

func replaceParagraphs(md string) string {
	var result strings.Builder

	htmlTags := []string{"<h", "<ul>", "<li>", "</ul>"}

	lines := strings.Split(md, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if stringHasPrefixes(line, htmlTags) {
			result.WriteString(line)
			continue
		}

		result.WriteString("<p>" + line + "</p>")
	}

	return result.String()
}

func stringHasPrefixes(s string, prs []string) bool {
	for _, pr := range prs {
		if strings.HasPrefix(s, pr) {
			return true
		}
	}

	return false
}