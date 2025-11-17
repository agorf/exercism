package house

import "strings"

var phrases = []string{
	"malt\nthat lay in the",
	"rat\nthat ate the",
	"cat\nthat killed the",
	"dog\nthat worried the",
	"cow with the crumpled horn\nthat tossed the",
	"maiden all forlorn\nthat milked the",
	"man all tattered and torn\nthat kissed the",
	"priest all shaven and shorn\nthat married the",
	"rooster that crowed in the morn\nthat woke the",
	"farmer sowing his corn\nthat kept the",
	"horse and the hound and the horn\nthat belonged to the",
}

func innerPhrase(v int) string {
	if v == 0 {
		return ""
	}
	return phrases[v-1] + " " + innerPhrase(v-1)
}

func Verse(v int) string {
	return "This is the " + innerPhrase(v-1) + "house that Jack built."
}

func Song() string {
	var verses []string
	for v := 1; v <= 12; v++ {
		verses = append(verses, Verse(v))
	}
	return strings.Join(verses, "\n\n")
}
