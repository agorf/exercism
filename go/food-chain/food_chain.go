package foodchain

import (
	"fmt"
	"strings"
)

var animals = [...]string{
	"", // Verses are 1-index based
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

var phraseByAnimal = map[string]string{
	"spider": "It wriggled and jiggled and tickled inside her.",
	"bird":   "How absurd to swallow a bird!",
	"cat":    "Imagine that, to swallow a cat!",
	"dog":    "What a hog, to swallow a dog!",
	"goat":   "Just opened her throat and swallowed a goat!",
	"cow":    "I don't know how she swallowed a cow!",
}

func Verse(v int) string {
	var lines []string
	animal := animals[v]
	lines = append(lines, fmt.Sprintf("I know an old lady who swallowed a %s.", animal))
	if animal == "horse" {
		lines = append(lines, "She's dead, of course!")
	} else {
		if phrase, exists := phraseByAnimal[animal]; exists {
			lines = append(lines, phrase)
		}
		for i := v; i >= 2; i-- {
			caughtAnimal := animals[i-1]
			if caughtAnimal == "spider" {
				caughtAnimal = "spider that wriggled and jiggled and tickled inside her"
			}
			lines = append(lines, fmt.Sprintf("She swallowed the %s to catch the %s.", animals[i], caughtAnimal))
		}
		lines = append(lines, "I don't know why she swallowed the fly. Perhaps she'll die.")
	}
	return strings.Join(lines, "\n")
}

func Verses(start, end int) string {
	var verses []string
	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
