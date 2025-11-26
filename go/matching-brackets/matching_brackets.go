package brackets

func Bracket(input string) bool {
	var stack []byte
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '[':
			stack = append(stack, ']') // Push
		case '{':
			stack = append(stack, '}') // Push
		case '(':
			stack = append(stack, ')') // Push
		case ']', '}', ')':
			if len(stack) == 0 {
				return false
			}
			b := stack[len(stack)-1]
			if b != input[i] {
				return false
			}
			stack = stack[:len(stack)-1] // Pop
		}
	}
	return len(stack) == 0
}
