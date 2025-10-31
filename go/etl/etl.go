package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	transormed := make(map[string]int)
	for points, letters := range in {
		for _, letter := range letters {
			transormed[strings.ToLower(letter)] = points
		}
	}
	return transormed
}
