package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	proverb := make([]string, 0, len(rhyme))
	if len(rhyme) == 0 {
		return proverb
	}
	for i := 1; i < len(rhyme); i++ {
		w1 := rhyme[i-1]
		w2 := rhyme[i]
		line := fmt.Sprintf("For want of a %s the %s was lost.", w1, w2)
		proverb = append(proverb, line)
	}
	line := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	proverb = append(proverb, line)
	return proverb
}
