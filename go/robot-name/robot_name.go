package robotname

import (
	"errors"
	"math/rand"
)

type Robot struct {
	name string
}

var (
	availableNames = []string{}
	populatedNames = false
)

func populateNames() {
	name := make([]rune, 5)
	for i := range 26 {
		for j := range 26 {
			for k := range 10 {
				for l := range 10 {
					for m := range 10 {
						name[0] = rune('A' + i)
						name[1] = rune('A' + j)
						name[2] = rune('0' + k)
						name[3] = rune('0' + l)
						name[4] = rune('0' + m)
						availableNames = append(availableNames, string(name))
					}
				}
			}
		}
	}
	populatedNames = true
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if !populatedNames {
		populateNames()
	}
	n := len(availableNames)
	if n == 0 {
		return "", errors.New("namespace exhausted")
	}
	i := rand.Intn(n)
	r.name = availableNames[i]
	availableNames[i] = availableNames[n-1] // Remove i (1/2)
	availableNames = availableNames[:n-1]   // Remove i (2/2)
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
