package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	diagram  []string
	children []string
}

var plantsByEncoding = map[rune]string{
	'C': "clover",
	'G': "grass",
	'R': "radishes",
	'V': "violets",
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	seenChild := make(map[string]struct{})
	for _, child := range children {
		if _, exists := seenChild[child]; exists {
			return nil, errors.New("found duplicate child")
		}
		seenChild[child] = struct{}{} // Seen
	}
	for _, cup := range diagram {
		if cup == '\n' {
			continue
		}
		if _, exists := plantsByEncoding[cup]; !exists {
			return nil, errors.New("invalid cup code")
		}
	}
	diagramRows := strings.Split(diagram, "\n")
	if len(diagramRows) != 3 {
		return nil, errors.New("wrong diagram format")
	}
	diagramRows = diagramRows[1:]
	if len(diagramRows[0]) != len(diagramRows[1]) {
		return nil, errors.New("uneven rows")
	}
	if len(diagramRows[0]) != len(children)*2 {
		return nil, errors.New("wrong number of columns")
	}
	return &Garden{
		diagram:  diagramRows,
		children: children,
	}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	sortedChildren := append([]string(nil), g.children...)
	sort.Strings(sortedChildren)
	childIndex := -1
	for i, c := range sortedChildren {
		if c == child {
			childIndex = i * 2
			break
		}
	}
	if childIndex == -1 {
		return []string{}, false
	}
	var plants []string
	firstRow := []rune(g.diagram[0])
	plants = append(plants, plantsByEncoding[firstRow[childIndex]])
	plants = append(plants, plantsByEncoding[firstRow[childIndex+1]])
	secondRow := []rune(g.diagram[1])
	plants = append(plants, plantsByEncoding[secondRow[childIndex]])
	plants = append(plants, plantsByEncoding[secondRow[childIndex+1]])
	return plants, true
}
