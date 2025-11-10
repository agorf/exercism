package flatten

func Flatten(nested any) any {
	output := []any{}
	for _, elem := range nested.([]any) {
		switch v := elem.(type) {
		case int:
			output = append(output, v)
		case []any:
			output = append(output, Flatten(v).([]any)...)
		case nil:
			// Ignore
		default:
			panic("should never happen")
		}
	}
	return output
}
