package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	matched, err := regexp.MatchString(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`, text)
	return err == nil && matched
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-~*=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`^[^"]*"[^"]*(?i:password)[^"]*"[^"]*$`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	for i, line := range lines {
		submatch := re.FindStringSubmatch(line)
		if len(submatch) == 2 { // User name matched
			lines[i] = "[USR] " + submatch[1] + " " + line
		}
	}
	return lines
}
