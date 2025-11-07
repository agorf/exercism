package wordy

import (
	"strconv"
	"strings"
)

const (
	questionPrefix = "What is "
	questionSuffix = "?"
	stateInitial   = iota
	statePlus
	stateMinus
	stateMultiplied
	stateMultipliedBy
	stateDivided
	stateDividedBy
	stateNumber
)

func Answer(question string) (int, bool) {
	if !strings.HasPrefix(question, questionPrefix) {
		return 0, false
	}
	if !strings.HasSuffix(question, questionSuffix) {
		return 0, false
	}
	question = strings.TrimPrefix(question, questionPrefix)
	question = strings.TrimSuffix(question, questionSuffix)
	parseNumber := func(s string) (int, bool) {
		num, err := strconv.Atoi(s)
		if err != nil {
			return 0, false
		}
		return num, true
	}
	words := strings.Fields(question)
	state := stateInitial
	result := 0
	for _, word := range words {
		switch state {
		case stateInitial:
			var ok bool
			result, ok = parseNumber(word)
			if !ok {
				return 0, false
			}
			state = stateNumber
		case stateNumber:
			switch word {
			case "plus":
				state = statePlus
			case "minus":
				state = stateMinus
			case "multiplied":
				state = stateMultiplied
			case "divided":
				state = stateDivided
			default:
				return 0, false
			}
		case statePlus:
			num, ok := parseNumber(word)
			if !ok {
				return 0, false
			}
			result += num
			state = stateNumber
		case stateMinus:
			num, ok := parseNumber(word)
			if !ok {
				return 0, false
			}
			result -= num
			state = stateNumber
		case stateMultiplied:
			if word != "by" {
				return 0, false
			}
			state = stateMultipliedBy
		case stateMultipliedBy:
			num, ok := parseNumber(word)
			if !ok {
				return 0, false
			}
			result *= num
			state = stateNumber
		case stateDivided:
			if word != "by" {
				return 0, false
			}
			state = stateDividedBy
		case stateDividedBy:
			num, ok := parseNumber(word)
			if !ok || num == 0 {
				return 0, false
			}
			result /= num
			state = stateNumber
		default:
			return 0, false
		}
	}
	if state != stateNumber {
		return 0, false
	}
	return result, true
}
