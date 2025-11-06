package twelve

import (
	"fmt"
	"strings"
)

var dayPresents = [...]string{
	"",
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

var numbers = [...]string{
	"",
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

const (
	verseFmt  = "On the %s day of Christmas my true love gave to me: %s."
	totalDays = 12
)

func presentsString(i int) string {
	if i == 1 {
		return dayPresents[1]
	}
	presents := make([]string, 0, i)
	for j := i; j >= 1; j-- {
		present := dayPresents[j]
		if j == 1 {
			present = "and " + present
		}
		presents = append(presents, present)
	}
	return strings.Join(presents, ", ")
}

func Verse(i int) string {
	return fmt.Sprintf(verseFmt, numbers[i], presentsString(i))
}

func Song() string {
	verses := make([]string, totalDays)
	for i := range verses {
		verses[i] = Verse(i + 1)
	}
	return strings.Join(verses, "\n")
}
