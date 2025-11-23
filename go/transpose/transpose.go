package transpose

import "fmt"

func padInput(input []string) []string {
	// Find longest row and its index
	padLen := 0
	padRow := 0
	for i, row := range input {
		rowLen := len(row)
		if padLen < rowLen {
			padLen = rowLen
			padRow = i
		}
	}
	numRows := len(input)
	var paddedInput []string
	for i, row := range input {
		rowPadLen := len(row)
		// Pad row until longest row to match it
		if i < padRow {
			rowPadLen = max(rowPadLen, padLen)
		} else if i < numRows-1 { // Pad to row below, if it exists
			rowPadLen = max(rowPadLen, len(input[i+1]))
		}
		paddedRow := fmt.Sprintf("%-*s", rowPadLen, row)
		paddedInput = append(paddedInput, paddedRow)
	}
	return paddedInput
}

func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}
	paddedInput := padInput(input)
	numCols := len(paddedInput[0])
	numRows := len(paddedInput)
	var transposedInput []string
	for j := range numCols {
		var row []byte
		for i := range numRows {
			if j < len(paddedInput[i]) {
				row = append(row, paddedInput[i][j])
			}
		}
		transposedInput = append(transposedInput, string(row))
	}
	return transposedInput
}
