package house

var phrases = []string{
	"the house that Jack built",
	"the malt\nthat lay in",
	"the rat\nthat ate",
	"the cat\nthat killed",
	"the dog\nthat worried",
	"the cow with the crumpled horn\nthat tossed",
	"the maiden all forlorn\nthat milked",
	"the man all tattered and torn\nthat kissed",
	"the priest all shaven and shorn\nthat married",
	"the rooster that crowed in the morn\nthat woke",
	"the farmer sowing his corn\nthat kept",
	"the horse and the hound and the horn\nthat belonged to",
}

func Song() string {
	var res string
	len := len(phrases)
	for i := 0; i < len; i++ {
		res += Verse(i + 1)
		if i != len-1 {
			res += "\n\n"
		}
	}
	return res
}

func Verse(index int) string {
	var res string
	res += "This is " + phrases[index-1]
	for i := index - 2; i >= 0; i-- {
		res += " " + phrases[i]
	}
	return res + "."
}
