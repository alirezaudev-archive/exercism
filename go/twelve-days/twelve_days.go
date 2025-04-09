package twelve

import "strings"

var days = []string{
	"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}

var gifts = []string{
	"a Partridge in a Pear Tree.",
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

func Verse(i int) string {
	prefix := "On the " + days[i-1] + " day of Christmas my true love gave to me: "
	parts := make([]string, 0, i)

	for j := i - 1; j >= 0; j-- {
		if j == 0 && i != 1 {
			parts = append(parts, "and "+gifts[j])
		} else {
			parts = append(parts, gifts[j])
		}
	}

	return prefix + strings.Join(parts, ", ")
}

func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n")
}
