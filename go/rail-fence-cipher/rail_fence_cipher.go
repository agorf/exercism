package railfence

type processFunc func(row, col int)

func buildFence(rows, cols int) [][]byte {
	var fence [][]byte
	for range rows {
		var row []byte
		for range cols {
			row = append(row, '.')
		}
		fence = append(fence, row)
	}
	return fence
}

func zigzagProcess(rows, cols int, process processFunc) {
	i, j := 0, 0
	inc := true
	for {
		process(i, j)
		if inc {
			i++
			if i == rows-1 {
				inc = false
			}
		} else {
			i--
			if i == 0 {
				inc = true
			}
		}
		j++
		if j == cols {
			break
		}
	}
}

func rowProcess(rows, cols int, process processFunc) {
	for i := range rows {
		for j := range cols {
			process(i, j)
		}
	}
}

func Encode(message string, rails int) string {
	fence := buildFence(rails, len(message))
	// Assign message into fence
	zigzagProcess(rails, len(message), func(i, j int) {
		fence[i][j] = message[j]
	})
	// Collect encoded message from fence
	var encoded []byte
	rowProcess(rails, len(message), func(i, j int) {
		b := fence[i][j]
		if b != '.' {
			encoded = append(encoded, b)
		}
	})
	return string(encoded)
}

func Decode(message string, rails int) string {
	fence := buildFence(rails, len(message))
	// Assign ? into fence to mark message positions
	zigzagProcess(rails, len(message), func(i, j int) {
		fence[i][j] = '?'
	})
	// Assign encoded message into fence
	k := 0
	rowProcess(rails, len(message), func(i, j int) {
		if fence[i][j] == '?' {
			fence[i][j] = message[k]
			k++
		}
	})
	// Collect decoded message from fence
	var decoded []byte
	zigzagProcess(rails, len(message), func(i, j int) {
		decoded = append(decoded, fence[i][j])
	})
	return string(decoded)
}
