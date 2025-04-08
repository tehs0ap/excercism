package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`).MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[~*-=]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	regex := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, line := range lines {
		if regex.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d*`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	taggedLines := make([]string, len(lines))
	regex := regexp.MustCompile(`User\s+([[:alnum:]]+)\s+`)

	for index, line := range lines {
		submatch := regex.FindStringSubmatch(line)
		if submatch != nil {
			taggedLines[index] = fmt.Sprintf("[USR] %s %s", submatch[1], line)
		} else {
			taggedLines[index] = line
		}
	}
	return taggedLines
}
