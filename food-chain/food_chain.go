package foodchain

import (
	"fmt"
)

type Pair struct {
	kind, desc string
}

var pairs = []Pair{
	{
		"",
		"",
	},
	{
		"fly",
		"I don't know why she swallowed the fly. Perhaps she'll die.",
	},
	{
		"spider",
		"It wriggled and jiggled and tickled inside her.",
	},
	{
		"bird",
		"How absurd to swallow a bird!",
	},
	{
		"cat",
		"Imagine that, to swallow a cat!",
	},
	{
		"dog",
		"What a hog, to swallow a dog!",
	},
	{
		"goat",
		"Just opened her throat and swallowed a goat!",
	},
	{
		"cow",
		"I don't know how she swallowed a cow!",
	},
}

func Verse(n int) (res string) {
	if n == 1 {
		res += "I know an old lady who swallowed a fly.\nI don't know why she swallowed the fly. Perhaps she'll die."
		return
	}

	if n == 8 {
		res += "I know an old lady who swallowed a horse.\nShe's dead, of course!"
		return
	}

	res += fmt.Sprintf("I know an old lady who swallowed a %s.\n%s\n", pairs[n].kind, pairs[n].desc)

	for i := n - 1; i > 0; i-- {
		if pairs[i].kind == "spider" {
			res += fmt.Sprintf("She swallowed the %s to catch the %s that wriggled and jiggled and tickled inside her.\n", pairs[i+1].kind, pairs[i].kind)
		} else {
			res += fmt.Sprintf("She swallowed the %s to catch the %s.\n", pairs[i+1].kind, pairs[i].kind)
		}
	}

	if n != 1 {
		res += "I don't know why she swallowed the fly. Perhaps she'll die."
	}

	return
}

func Verses(start, end int) (res string) {

	for i := start; i <= end; i++ {
		res += Verse(i)
		if i != end {
			res += "\n\n"
		}
	}

	return
}

func Song() string {
	return Verses(1, 8)
}
