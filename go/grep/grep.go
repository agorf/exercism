package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func lineMatchesPattern(line, pattern string, flagsMap map[string]bool) bool {
	if flagsMap["-i"] {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}
	if flagsMap["-x"] {
		return line == pattern
	}
	return strings.Contains(line, pattern)
}

func Search(pattern string, flags, filenames []string) []string {
	flagsMap := map[string]bool{}
	for _, flag := range flags {
		flagsMap[flag] = true
	}
	var matches []string
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		sc := bufio.NewScanner(f)
		lineNo := 1
		for sc.Scan() {
			line := sc.Text()
			matched := lineMatchesPattern(line, pattern, flagsMap)
			if flagsMap["-v"] {
				matched = !matched
			}
			if matched {
				if flagsMap["-l"] {
					matches = append(matches, filename)
					break
				}
				var formatParts []string
				var args []any
				if len(filenames) > 1 {
					formatParts = append(formatParts, "%s")
					args = append(args, filename)
				}
				if flagsMap["-n"] {
					formatParts = append(formatParts, "%d")
					args = append(args, lineNo)
				}
				formatParts = append(formatParts, "%s")
				args = append(args, line)
				format := strings.Join(formatParts, ":")
				matches = append(matches, fmt.Sprintf(format, args...))
			}
			lineNo++
		}
		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return matches
}
