package twelve
import (
	"bytes"
	"fmt"
)

type wrapper struct {
	dayth string
	gift string
}

var sink = map[int]wrapper{
	1 : wrapper{"first", "a Partridge in a Pear Tree"},
	2 : wrapper{"second", "two Turtle Doves"},
	3 : wrapper{"third", "three French Hens"},
	4 : wrapper{"fourth", "four Calling Birds"},
	5 : wrapper{"fifth", "five Gold Rings"},
	6 : wrapper{"sixth", "six Geese-a-Laying"},
	7 : wrapper{"seventh", "seven Swans-a-Swimming"},
	8 : wrapper{"eighth", "eight Maids-a-Milking"},
	9 : wrapper{"ninth", "nine Ladies Dancing"},
	10 : wrapper{"tenth", "ten Lords-a-Leaping"},
	11 : wrapper{"eleventh", "eleven Pipers Piping"},
	12 : wrapper{"twelfth", "twelve Drummers Drumming"},
}

func Verse(day int) string {
	return format(day)
}

func Song() string {
	var buf bytes.Buffer
	for i := 1; i <= 11; i++ {
		fmt.Fprintf(&buf, "%s\n", format(i))
	}
	fmt.Fprintf(&buf, "%s", format(12))
	return buf.String()
}

func format(day int) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "On the %s day of Christmas my true love gave to me: %s", sink[day].dayth, sink[day].gift)

	for i := day - 1; i >= 1; i-- {
		if i == 1{
			fmt.Fprintf(&buf, ", and %s",sink[i].gift)
		}else{
			fmt.Fprintf(&buf, ", %s",sink[i].gift)
		}
	}
	buf.WriteByte('.')
	return buf.String()
}